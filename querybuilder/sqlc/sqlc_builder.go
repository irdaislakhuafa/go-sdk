package sqlc

import (
	"fmt"
	"strconv"
	"strings"
)

var _ BuilderInterface = &Builder{}

// Determine condition on "SELECT x FROM x WHERE {{ this is your expression }}".
//
// Where("id = ?", 1) - If no logic is set, then will use "AND" by default
//
// Where("OR email = ?", "x@gmail.com") - If logic has set, then will use it instead of "AND" (default)
//
// Where("is_deleted = $1", 0)
//
// Where("AND id IN (SELECT user_id FROM user_roles WHERE role_id = $1)", 2)
func (b *Builder) Where(expression string, args ...any) BuilderInterface {
	b.filters = append(b.filters, filter{
		expression: expression,
		hasLogic: strings.HasPrefix(expression, "and") ||
			strings.HasPrefix(expression, "or") ||
			strings.HasPrefix(expression, "AND") ||
			strings.HasPrefix(expression, "OR"),
		args: args,
	})
	return b
}

// Like method Where() but will add "OR" logic to your expression
//
// Then your query will look like "OR {{ your expression here }}"
//
// Example:
//
// Or("name LIKE ?", "%me%")
func (b *Builder) Or(column string, args ...any) BuilderInterface {
	b.filters = append(b.filters, filter{
		expression: fmt.Sprintf("OR %s", column),
		hasLogic:   true,
		args:       args,
	})
	return b
}

// Like method Where() but will add "AND" logic to your expression
//
// Then your query will look like "AND {{ your expression here }}"
//
// Example:
//
// And("name LIKE ?", "%me%")
func (b *Builder) And(column string, args ...any) BuilderInterface {
	b.filters = append(b.filters, filter{
		expression: fmt.Sprintf("AND %s", column),
		hasLogic:   true,
		args:       args,
	})
	return b
}

// Is equal with Where("id IN (?,?,?,...)", args...)
//
// Example:
//
// In("id", 1,2,3)
//
// In("OR id", 1,2,3)
func (b *Builder) In(column string, args ...any) BuilderInterface {
	inExpression := strings.Repeat("?,", len(args))
	inExpression = inExpression[:len(inExpression)-1]
	column += fmt.Sprintf(" IN (%s)", inExpression)
	return b.Where(column, args...)
}

// Order columns on SELECT query. Your query will like "SELECT x FROM x WHERE x ORDER BY {{ cols }}"
//
// Example:
//
// Order("id DESC")
//
// Order("id, age DESC")
func (b *Builder) Order(cols string, args ...any) BuilderInterface {
	b.order = order{
		expression: cols,
		args:       args,
	}
	return b
}

// Limit for rows that returned on SELECT query.
//
// Example:
//
// Limit(10)
func (b *Builder) Limit(limit int) BuilderInterface {
	b.limit = &limit
	return b
}

// Offset for rows that returned on SELECT query.
//
// Example:
//
// Offset(10)
func (b *Builder) Offset(offset int) BuilderInterface {
	b.offset = &offset
	return b
}

// Build or compile your queries
func (b *Builder) Build(query string, args ...any) (string, []any) {
	sb := strings.Builder{}
	sb.WriteString(query)
	sb.WriteByte('\n')

	isContainWhere := strings.Contains(strings.ToLower(query), "where")
	for i, f := range b.filters {
		if i == 0 && !isContainWhere {
			sb.WriteString("WHERE ")
		} else {
			if !f.hasLogic {
				sb.WriteString("AND ")
			}
		}

		if f.hasLogic {
			if i == 0 && isContainWhere {
				sb.WriteString("AND 1 = 1\n")
			} else if i == 0 && !isContainWhere {
				sb.WriteString("1 = 1\n")
			}
			sb.WriteString(f.expression)
			sb.WriteByte('\n')
		} else {
			sb.WriteByte('(')
			sb.WriteString(f.expression)
			sb.WriteString(")\n")
		}

		args = append(args, f.args...)
	}

	if b.order.expression != "" {
		sb.WriteString("ORDER BY ")
		sb.WriteString(b.order.expression)
		sb.WriteByte('\n')
		args = append(args, b.order.args...)
	}

	if b.limit != nil {
		sb.WriteString("LIMIT ")
		sb.WriteString(strconv.Itoa(*b.limit))
		sb.WriteByte('\n')
	}

	if b.offset != nil {
		sb.WriteString("OFFSET ")
		sb.WriteString(strconv.Itoa(*b.offset))
		sb.WriteByte('\n')
	}

	return sb.String(), args
}

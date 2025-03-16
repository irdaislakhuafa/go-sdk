# Overview

[SQLC](https://github.com/sqlc-dev/sqlc) is a great tool to use raw sql query with golang. But currently `sqlc` lacks of dynamic query support.

It's currently only support static query that compiled to golang code, but sometimes we need dynamic query for condition/order/limit/etc. So i write this code as wrapper for sqlc for dynamic query.

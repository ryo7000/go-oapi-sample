go-oapi-sample
=====

libraries
-----

* slog
* echo
* gorm
* wire
* oapi-codegen
* kin-openapi
* ozzo-validation

setup
-----

* Install Docker and Compose plugin
* Install task runner

  ```
  go install github.com/go-task/task/v3/cmd/task@latest
  ```

start dev
-----

1. Start docker
2. Start postgresql by `task db:start`
3. Create database if necessary by `task db:create`
4. Migrate database if necessary by `task db:migrate:up`
5. Start server with hot reload by `task dev`


show all tasks
-----

```
task
```

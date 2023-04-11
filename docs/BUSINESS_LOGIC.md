# ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which caching solution your choose_ or any third-party things.

## Modules

- `.app/modules/[module]/dto` folder for all dto implementations
- `.app/modules/[module]/entities` folder for describe business models of your project
- `.app/modules/[module]/repositories` folder for describe queries for models of your project
- `.app/modules/[module]/services` folder for describe business implementation for your controllers
- `.app/modules/[module]/module_controller.go` file for functional controllers (used in routes)
- `.app/modules/[module]/module_routes.go` file for all routes about current module

## Middleware

- `./app/middleware` folder for add middleware (Fiber and yours)

Files must be named like `name_middleware.go`

## Permissions

- `./app/permissions` folder for describe `credentials, roles` of your project

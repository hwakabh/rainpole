# Changelog

## [0.3.0](https://github.com/hwakabh/rainpole/compare/v0.2.0...v0.3.0) (2025-03-10)


### Features

* enabled to fetch fixtured data from SQLite3 database. ([d58232c](https://github.com/hwakabh/rainpole/commit/d58232cf0c145ce597a89978b19b55f7d64831e8))


### Bug Fixes

* **ci:** error access with test database. ([5f9cd68](https://github.com/hwakabh/rainpole/commit/5f9cd68411cf441de383a858ce25aa38e4c7a23e))
* **ci:** error access with test database. ([77c06b8](https://github.com/hwakabh/rainpole/commit/77c06b886563880248b1c563c4165353d25f542b))


### Documentation

* **build:** replaced commands with make. ([245f01a](https://github.com/hwakabh/rainpole/commit/245f01a4e439319f8416c147f99f52c8e51c0e29))


### Miscellaneous Chores

* added initial snippets to create clean database. ([504c005](https://github.com/hwakabh/rainpole/commit/504c0056d62e6a5909954f0d94e1c5d1fc61334b))
* **deps:** update dependency go to v1.24.1 ([06a3e1d](https://github.com/hwakabh/rainpole/commit/06a3e1df6ccd80e9537c22c3e6782972bec6bb29))
* renamed asstes directory from web to public. ([d4afe5a](https://github.com/hwakabh/rainpole/commit/d4afe5a1ea7448c04f16368b65bb913af16ad872))


### Code Refactoring

* **build:** replaced scripts with Makefile. ([169033d](https://github.com/hwakabh/rainpole/commit/169033db1a9b1056532a67d850907821d62902a1))
* **build:** updated Makefile. ([ba6c193](https://github.com/hwakabh/rainpole/commit/ba6c1932af520947fa49fd073c065518871a63e8))

## [0.2.0](https://github.com/hwakabh/rainpole/compare/v0.1.0...v0.2.0) (2025-02-27)


### Features

* **api:** implemented random uuid endpoints. ([efd3312](https://github.com/hwakabh/rainpole/commit/efd33128ee7a01ecd7810535ed3723dc401682d4))
* implemented calver-based tags. ([c2f5108](https://github.com/hwakabh/rainpole/commit/c2f5108011f6eaee59ebc0794ae5958cc27f1417))
* implemented validations with UUID custom-types. ([e33d799](https://github.com/hwakabh/rainpole/commit/e33d799246c39a6e1694ae10b152a1972afeb675))


### Bug Fixes

* **build:** binaries path. ([6dd18f2](https://github.com/hwakabh/rainpole/commit/6dd18f21a2fd01262a6f33b604309b927dfb99e0))
* **ci:** added missing envars to push step. ([dcbbeab](https://github.com/hwakabh/rainpole/commit/dcbbeab0474835b7fb8b66d84d2191247ea970d8))
* **ci:** error on Cloud Run with blank values. ([b5e0309](https://github.com/hwakabh/rainpole/commit/b5e0309e1aa503e8795863bf01e13d226d751982))
* **ci:** path error on Test CI. ([747ec70](https://github.com/hwakabh/rainpole/commit/747ec7015d529c2add96230471abf87d1eb34a8e))
* **ci:** syntax referencing calver-based tag. ([f4da078](https://github.com/hwakabh/rainpole/commit/f4da0789aa05368bf349b67a0173ccfb258bfff5))


### Reverts

* CMD to ENTRYPOINT for invoking binaries. ([c5b8c97](https://github.com/hwakabh/rainpole/commit/c5b8c97e5489668e8de2f211760bd85631b829ca))


### Documentation

* added notes for build. ([fa8c956](https://github.com/hwakabh/rainpole/commit/fa8c9567bbf240d0f182fc52784c96c8a33d7547))


### Miscellaneous Chores

* added packer validation on Test CI. ([8df8587](https://github.com/hwakabh/rainpole/commit/8df858769d88cd0e8cbc2e6638aaa7a3040f3ae5))
* **build:** added build scripts. ([2ab2f43](https://github.com/hwakabh/rainpole/commit/2ab2f430350212331c1ee07658e7ee4ec83fcb71))
* **builds:** added single quotes ([1b0c96a](https://github.com/hwakabh/rainpole/commit/1b0c96ac1b885d67890e6a0784a78fed5f6c9abb))
* **builds:** enabled service-account in build configs. ([fe8ebcd](https://github.com/hwakabh/rainpole/commit/fe8ebcd6636db31f9333a95f5d54d60a65b30430))
* **builds:** enabled variables substitutions. ([9013f7a](https://github.com/hwakabh/rainpole/commit/9013f7ab66dd2489d7c9f4664746d180fc356e12))
* **builds:** fixed proper image name. ([e970a4a](https://github.com/hwakabh/rainpole/commit/e970a4ab4557f94a42c493ec300bf97282e7596b))
* **builds:** updated docs for triggers. ([d841892](https://github.com/hwakabh/rainpole/commit/d8418925da393b7a2c9fb39329e687aa2e4dc514))
* **builds:** updated docs for triggers. ([70369bc](https://github.com/hwakabh/rainpole/commit/70369bc5e28c534d5be62294882b054638a2bad5))
* **builds:** updated docs for triggers. ([e0d27a9](https://github.com/hwakabh/rainpole/commit/e0d27a9a09aa40a33c09d950894b2ad194344460))
* **build:** updated startups from ENTRYPOINT to CMD for distroless. ([6c4cd86](https://github.com/hwakabh/rainpole/commit/6c4cd863f4eafd14edb96aa33da1e98ea1b3f5d0))
* **ci:** enhanced release-please configs. ([28cdbb6](https://github.com/hwakabh/rainpole/commit/28cdbb6266307659c2227eb8e02b686615919d29))
* **ci:** fixed workflows for utilizing packer. ([88b6797](https://github.com/hwakabh/rainpole/commit/88b6797d2edc924aca9fa8d9e44c9fcab67042a7))
* **deps:** update all non-major dependencies to v1.23.6 ([6160d96](https://github.com/hwakabh/rainpole/commit/6160d9679e090126e6265020f8e4f78ba9191511))
* **deps:** update all non-major dependencies to v1.24.0 ([08c5868](https://github.com/hwakabh/rainpole/commit/08c5868927f385421d44aee00440eb300b0e5f3a))
* **deps:** update docker/build-push-action action to v6.13.0 ([6b76f0f](https://github.com/hwakabh/rainpole/commit/6b76f0f1fa40e0fc7bf446fb9abd89c511a11503))
* **deps:** update docker/build-push-action action to v6.14.0 ([ba709c6](https://github.com/hwakabh/rainpole/commit/ba709c6f00b3d8541cb27da9370c3310f9e19487))
* finalized triggers. ([7cd1504](https://github.com/hwakabh/rainpole/commit/7cd1504b3fea50cd18d309b39e376f6bc3acbc2d))
* fixed dependencies between steps. ([59182a3](https://github.com/hwakabh/rainpole/commit/59182a3cfe2c2b45eb7075fb25b03276c0bcfa49))
* fixed package name in the module. ([2245e6b](https://github.com/hwakabh/rainpole/commit/2245e6b3339464e470fdc15e096a7671161bdaf3))
* implemented serving static file in root endpoint. ([d4d4098](https://github.com/hwakabh/rainpole/commit/d4d4098b1c13fff1a506f0fddec34486df11bc99))
* made minor changes in setup-go ([06db48f](https://github.com/hwakabh/rainpole/commit/06db48f50e3f1700c3ff686b48392bd730a78b6f))
* removed unused Cloud Build configs. ([14f14fe](https://github.com/hwakabh/rainpole/commit/14f14fe902af90d9b1da802d93d2cde449c25174))
* renamed packer templates. ([74e344c](https://github.com/hwakabh/rainpole/commit/74e344ced5f93bd8bd9d3d34901a55631eb4f1c0))
* replaced base distroless images. ([d7837ed](https://github.com/hwakabh/rainpole/commit/d7837edfd2b6a696c45a9296b3548483fd73fe0a))
* replaced base distroless images. ([bce32fc](https://github.com/hwakabh/rainpole/commit/bce32fcf2270bd987eec765879a86ee6413f65f2))
* separated functions per endpoints. ([f27d129](https://github.com/hwakabh/rainpole/commit/f27d1296278fbeec837c6281216579bba7685f63))
* separated test files. ([567c15c](https://github.com/hwakabh/rainpole/commit/567c15c461fa4dcdfa58955b4f490b5051e4bc0f))
* updated editorconfigs for HCL files. ([0951029](https://github.com/hwakabh/rainpole/commit/095102953268a8ac4b24f1a54d745f28ae3eeca3))
* updated pre-defined variable name. ([4ce8da2](https://github.com/hwakabh/rainpole/commit/4ce8da23bbdcbff61b0cd5cfc589acd762b52e20))


### Code Refactoring

* **ci:** replaced Dockerfiles with packer build. ([a181a87](https://github.com/hwakabh/rainpole/commit/a181a87936ea0d98a23014b9a54bf243717788f9))


### Tests

* added subtestings for default routes. ([82c6256](https://github.com/hwakabh/rainpole/commit/82c6256fe184245dd59823aed8f85be304096a63))
* **api:** added tests for REST-API endpoints. ([a87797b](https://github.com/hwakabh/rainpole/commit/a87797bb0c86cf9420e5fe5a1fe20d473859ba03))
* enabled to check response body with root/health endpoint. ([b902f16](https://github.com/hwakabh/rainpole/commit/b902f16077342134c88cede49f06a4a287885645))

## [0.1.0](https://github.com/hwakabh/rainpole/compare/v0.0.1...v0.1.0) (2025-01-23)


### Features

* added MVP with net/http server, handler, and multiplexer. ([b35094b](https://github.com/hwakabh/rainpole/commit/b35094b2029d5fc38f0c876aa9ef9c45803beb21))
* **ci:** finalized CI triggers. ([10d5ba2](https://github.com/hwakabh/rainpole/commit/10d5ba20751ac71384f9d34b6f5e25a6efb434fa))
* **ci:** implemented deploy pipeline to Cloud Run. ([54c130c](https://github.com/hwakabh/rainpole/commit/54c130c7e0fad4b3d9d9461801f0cfab285e920f))
* updated endpoints with JSON reponses. ([785b2ea](https://github.com/hwakabh/rainpole/commit/785b2ea95664eb855860d9570b47edf07f007ca8))
* updated service-account for WIF with Cloud Run. ([3ac5a80](https://github.com/hwakabh/rainpole/commit/3ac5a807ee1c9694fda73fe2c9531fa382785f1e))


### Bug Fixes

* **ci:** resolved build triggers with each PR push. ([97f9df2](https://github.com/hwakabh/rainpole/commit/97f9df2b67ea9dd086f0073644374f72421ed424))
* implemented RestMessage struct with same orders in each functions. ([a76ad8c](https://github.com/hwakabh/rainpole/commit/a76ad8c69b38cebf2b7ccb8bf46b72f2221af44e))
* known issues on Google Cloud endpoints. ([9bd5478](https://github.com/hwakabh/rainpole/commit/9bd54788cfafffd420a526835d4e4ab6940bcd28))

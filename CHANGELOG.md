# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

## [1.2.0](https://github.com/Tomut0/SimpleClansStats/compare/v1.1.0...v1.2.0) (2024-12-14)


### Features

* add abstract layer of tables (DataTable.vue) ([c55a43b](https://github.com/Tomut0/SimpleClansStats/commit/c55a43bbeba7217d6d6a23c34a0d4b1d66043f14))
* add clan modal and entity selector ([5d88210](https://github.com/Tomut0/SimpleClansStats/commit/5d88210fc77017e0fd438c8f8fe78dbecd637ceb))
* add explanation how to build Docker compose locally ([d9ab004](https://github.com/Tomut0/SimpleClansStats/commit/d9ab0047e9dff674cd9f12c8f33520a3be1c4948))
* add og tags and image ([db0a0ce](https://github.com/Tomut0/SimpleClansStats/commit/db0a0ce2c7ceb219ca0d6da0215ae4a30135ea08))
* consider db table prefix ([4585e90](https://github.com/Tomut0/SimpleClansStats/commit/4585e904ef2fa3ce376d2343e6ebc95e34c6b66d))
* draw clan banners on clan's model ([56fe779](https://github.com/Tomut0/SimpleClansStats/commit/56fe7798e1e9495599e3e0eb00440c681603c070))
* impl ClanPlayerFactory to generate fake clanplayers data ([d283966](https://github.com/Tomut0/SimpleClansStats/commit/d2839663047789c5d3a5a1118cc9222ff7a1896f))
* impl SCSSeeder to seed the database with fake data ([484f8b1](https://github.com/Tomut0/SimpleClansStats/commit/484f8b129f2bd9808c61b816f1afe4a5bb63bbc8))


### Bug Fixes

* avoid select same selector value ([95b337d](https://github.com/Tomut0/SimpleClansStats/commit/95b337d395fef0e69a96278d7739b32c3f1aa79d))
* count ally kills at kdr ([3bbf5db](https://github.com/Tomut0/SimpleClansStats/commit/3bbf5dbea566b7dfb429c9785d4e5445c6715851))
* members count interfere with the relationship ([6108742](https://github.com/Tomut0/SimpleClansStats/commit/6108742f0b6a521e228f842ffe4767f8fe1f97b2))
* move api routes to v1/api ([36f6aa6](https://github.com/Tomut0/SimpleClansStats/commit/36f6aa68bfe253d2993fae02067d229485ee1659))
* rename Dashboard page to Leaderboard ([790f20d](https://github.com/Tomut0/SimpleClansStats/commit/790f20d0b9b56ce1f1a9b9f98dc58228afbf449d))
* **ui:** charts width based on count ([1102f9d](https://github.com/Tomut0/SimpleClansStats/commit/1102f9d1b1073fcefedcd8276a1d3f3f95057be2))
* upgrade @heroicons/vue from 2.1.3 to 2.1.4 ([c1b1cef](https://github.com/Tomut0/SimpleClansStats/commit/c1b1ceff3fb7bb1864e1b7024acda001cf0baf22))
* upgrade @vueuse/motion from 2.2.3 to 2.2.4 ([a7155db](https://github.com/Tomut0/SimpleClansStats/commit/a7155db14d8b20d5a77c870cf5d52d637e3d559f))
* upgrade @vueuse/motion from 2.2.4 to 2.2.5 ([f5950e4](https://github.com/Tomut0/SimpleClansStats/commit/f5950e43de300fd090d726993d94e893f47796cc))
* upgrade @vueuse/motion from 2.2.5 to 2.2.6 ([38faaeb](https://github.com/Tomut0/SimpleClansStats/commit/38faaebb510af5634f8517d091c32ae2a5e19373))
* upgrade chart.js from 4.4.4 to 4.4.5 ([8e04c8f](https://github.com/Tomut0/SimpleClansStats/commit/8e04c8f0792d3f30b8d971c98761e68882eb5da1))
* upgrade multiple dependencies with Snyk ([91e7e3f](https://github.com/Tomut0/SimpleClansStats/commit/91e7e3fb4418506befa712a11692b391dcf00351))

## [1.1.0](https://github.com/Tomut0/SimpleClansStats/compare/v1.0.0...v1.1.0) (2024-07-21)


### Features

* add Dockerfile and docker compose ([bc673ea](https://github.com/Tomut0/SimpleClansStats/commit/bc673ea05e028b9d23100ed1832d87f8550a67cd))


### Bug Fixes

* env for kill weight is not allowed outside of config files ([2301abc](https://github.com/Tomut0/SimpleClansStats/commit/2301abc3258b54610e1ac9c6bdc462181315ec99))

## 1.0.0 (2024-06-23)


### Features

* add possibility to use sqlite as datasource ([3b4d23c](https://github.com/Tomut0/SimpleClansStats/commit/3b4d23ccccff84ae4a0737defb056a7b320f9e38))
* Dashboard page impl ([bcaef41](https://github.com/Tomut0/SimpleClansStats/commit/bcaef412be088d767719bc69ddf6d0e32e192a54))


### Bug Fixes

* avoid stackoverflow on translations ([1cd30c2](https://github.com/Tomut0/SimpleClansStats/commit/1cd30c2acde1b7697ee5a862449ccb489cce7b90))
* circular translation on empty key ([6a9bd74](https://github.com/Tomut0/SimpleClansStats/commit/6a9bd74f53ac35ce1085985298a009df52475e83))
* NPE on untranslated messages ([65bace8](https://github.com/Tomut0/SimpleClansStats/commit/65bace86504d22ce59dfe7bf6300ddcf97491428))
* **ui:** Translator should be above other elements ([dca67a4](https://github.com/Tomut0/SimpleClansStats/commit/dca67a440b378cda9bae7bb610effb52b52750a4))
* use eloquent relationship in the modals ([e878fb1](https://github.com/Tomut0/SimpleClansStats/commit/e878fb16f70756a0d9153023c6808f84cb33d0f9))

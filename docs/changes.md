

## [0.1.2](https://github.com/shoshta73/homehub/compare/0.1.1...0.1.2) (2024-10-31)

### Features

* **frontend/menubar:** logout functionality ([6063ad8](https://github.com/shoshta73/homehub/commit/6063ad809ce6b2dcaea3493dee1a4b760d873b48))
* **frontend/menubar:** validate return pattern ([a866c38](https://github.com/shoshta73/homehub/commit/a866c388a68dee12e982d73aae424730a162e651))
* **frontend:** paginate login view ([eec0121](https://github.com/shoshta73/homehub/commit/eec012144a568ec426de7b89f465af70dd7b024d))
* **frontend:** poll for avatar ([de672ed](https://github.com/shoshta73/homehub/commit/de672edbee32223b617cb532285f4a7a583a4a45))
* **frontend:** validate on protected routes ([a6d8d55](https://github.com/shoshta73/homehub/commit/a6d8d55e37aef72b1eccfc33e55b07b2720e9ead))
* **server:** implement email only login ([12afa3c](https://github.com/shoshta73/homehub/commit/12afa3c6551af1a195c65980b1ade9e50602bfed))
* **server:** implement username only login ([33d33e5](https://github.com/shoshta73/homehub/commit/33d33e5a6241678914535bf6313014b78dd7ba0c))
* **server:** implement validate handler ([9521452](https://github.com/shoshta73/homehub/commit/95214523a0240a704dc1eaa9022c74fcb2c1bb9a))
* **server:** logout handler ([b4fbe51](https://github.com/shoshta73/homehub/commit/b4fbe5118db8ecaad63b7aaa192e89aec3259691))
* **server:** Only 200 resp on /avatar ([23cf658](https://github.com/shoshta73/homehub/commit/23cf658ae64d1df5641d88fe95ad337faf5e45bf))
* **server:** prevent giving avatar url if image is not ready ([3da8e1a](https://github.com/shoshta73/homehub/commit/3da8e1a151ce4a4cc5f153b1f236738c393c46a5))
* **server:** simplify login logic ([fb6947f](https://github.com/shoshta73/homehub/commit/fb6947f2813ed57f4bd2d7c2b03b3a4c87ffb707))

### Performance Improvements

* **server:** improve generation of identicon ([b314c4a](https://github.com/shoshta73/homehub/commit/b314c4a08a0f411ff814ddb7699c6398f1978ce5))

### Miscellaneous Chores

* update changelog ([2b7fe07](https://github.com/shoshta73/homehub/commit/2b7fe07101658499c57bc16ed5b712c77359ecbb))

## [0.1.1](https://github.com/shoshta73/homehub/compare/0.1.0...0.1.1) (2024-10-30)

### Features

* **frontend:** guard api url from being undefined ([8708010](https://github.com/shoshta73/homehub/commit/870801040669ce15db987c05153105b3572d4a48))
* **frontend:** implement showing user stats ([6a3c6f7](https://github.com/shoshta73/homehub/commit/6a3c6f79025196b0bdc30a6233367c093714e0a4))
* **frontend:** send credentials to the /avatar and /pastebin/create ([4f039d1](https://github.com/shoshta73/homehub/commit/4f039d1185a1df2f215649937c47e424eb1e567f))
* **server:** pastebin stats handler ([8c80b7b](https://github.com/shoshta73/homehub/commit/8c80b7b0d0d928df8975e934464f2f1d5c29e268))
* **server:** prod/release mode ([8778939](https://github.com/shoshta73/homehub/commit/8778939e0382f202a0a4f83c64f74276d88b4c2f))

### Bug Fixes

* **server/models:** paste model now use time.Time for time stamping ([6e395ad](https://github.com/shoshta73/homehub/commit/6e395ad9e143021e82499a8d404a65d05b396283))
* **server:** Past id not being autoincrement ([a81ac51](https://github.com/shoshta73/homehub/commit/a81ac5180994bdbf5b3ecad6a02d3194736eecfa))

### Miscellaneous Chores

* **frontend:** added proxy config for development ([aaa2416](https://github.com/shoshta73/homehub/commit/aaa24160668b8fcb9a625967ba69c3facdf4318a))
* **frontend:** simplify config ([a057369](https://github.com/shoshta73/homehub/commit/a05736984d67565ddb163120e247a14a084db348))
* **server:** dev setup ([4e22436](https://github.com/shoshta73/homehub/commit/4e224365cea111418dcc534d37c127bbff8b0557))
* update changelog ([2a0477b](https://github.com/shoshta73/homehub/commit/2a0477bf0c334478998e848731ccc5f104ef431e))
* update Changelog ([ba46f6c](https://github.com/shoshta73/homehub/commit/ba46f6c046e591fd7f05bd075a9ac333618aeb7f))
* update go.mod file ([9a3d84c](https://github.com/shoshta73/homehub/commit/9a3d84c3401536dfb2271d1c1b10821891b7efce))

### Code Refactoring

* **server:** move handlers to their own module ([e0d51b4](https://github.com/shoshta73/homehub/commit/e0d51b455817c5d666ef6b04e450bd43cf3cd2e4))
* **server:** use default charmbracelet logger ([64e331f](https://github.com/shoshta73/homehub/commit/64e331f5c3f8cdb130aaf165214cde2718c19d5a))

### Build System

* update build configs to support building for prod ([43b6fdf](https://github.com/shoshta73/homehub/commit/43b6fdf60176b02d8b09fd3ff2ca5f0d4e311464))

## 0.1.0 (2024-10-28)

### Features

* **frontend/menubar:** implement displaying user avatar ([53212cc](https://github.com/shoshta73/homehub/commit/53212ccc36444b108f41e0bab3e62bbf5275e8b6))
* **frontend/navbar:** increase size of buttons and add login/register buttons ([04f431a](https://github.com/shoshta73/homehub/commit/04f431acef49447d9325bdbf64c0e3b232227324))
* **frontend/pastebin:** add fucntionality to the pastebin view ([443df81](https://github.com/shoshta73/homehub/commit/443df81d3fdef42b215abd93688c033a457c0c15))
* **frontend:** add input for paste title ([63f6fde](https://github.com/shoshta73/homehub/commit/63f6fdee6237d12ca3171f9f743cca5be0424d77))
* **frontend:** base for user home view ([47a2c54](https://github.com/shoshta73/homehub/commit/47a2c54c238a6186fb5f09770bf347250eab9def))
* **frontend:** dark-light mode toggle ([dbd1630](https://github.com/shoshta73/homehub/commit/dbd163035259101371f9501656da476757ed3088))
* **frontend:** functional login/register screens ([07d5ae1](https://github.com/shoshta73/homehub/commit/07d5ae121408cd113efa75571d84ef2fe4d495a7))
* **frontend:** implement simple pastebin view ([3cf8893](https://github.com/shoshta73/homehub/commit/3cf8893f1bb7feb8c56a1e2f56f408f133e2c1fa))
* **frontend:** login view ([5cdf5df](https://github.com/shoshta73/homehub/commit/5cdf5df35faae1065b0ce5d0ce07fafc5a32a6c9))
* **frontend:** registration view ([3287834](https://github.com/shoshta73/homehub/commit/328783475e87db030cb8bc98277101e0796fa628))
* **frontend:** use hashbrowser and use lazy import on register view ([1342cc7](https://github.com/shoshta73/homehub/commit/1342cc7390ab0afb939c1cc552b584539f89d937))
* **server/log:** create custom logger style ([3ff86d8](https://github.com/shoshta73/homehub/commit/3ff86d8a2ec3bbd2ee056a8f5a64f15d0c64523c))
* **server/log:** log with debug mode ([eaae7a4](https://github.com/shoshta73/homehub/commit/eaae7a499e25bc2e3465365ee73d26bb3fb172bb))
* **server/log:** report timestap with miliseconds ([af511ba](https://github.com/shoshta73/homehub/commit/af511ba255ca0065e05951f129b4cfd4f33c976d))
* **server/pastebin:** save to compressed file if contents to long ([0011e6b](https://github.com/shoshta73/homehub/commit/0011e6b5d27c4e6e8a713d6875a3180a48c59e51))
* **server:** create identicon ([98e1c38](https://github.com/shoshta73/homehub/commit/98e1c38cddc46e8f158e8a46e74fbf4bf258497a))
* **server:** create log module ([dfe5a66](https://github.com/shoshta73/homehub/commit/dfe5a66229d472bdfb49a26d34c627f10969ff7a))
* **server:** implement getting user avatar ([40a66a3](https://github.com/shoshta73/homehub/commit/40a66a33b592b8cc1aa98d4e050be23e941f330b))
* **server:** implement register handler ([182f339](https://github.com/shoshta73/homehub/commit/182f339e4fa218620e0dc30f17c3b8c9847a6fa3))
* **server:** implement sending cookie on register/login ([cc60304](https://github.com/shoshta73/homehub/commit/cc603044b9f19153e55a622a4f02503bf09e4e80))
* **server:** implement tracking pastebin stats for pastebin implementation ([d958062](https://github.com/shoshta73/homehub/commit/d95806293980ad841ee357257194e57bc23bdd34))
* **server:** logger preserves old logfiles ([8ef3f23](https://github.com/shoshta73/homehub/commit/8ef3f23ea466326a34c69b72750b079109307a81))
* **server:** multilogger now reports location ([9b141cf](https://github.com/shoshta73/homehub/commit/9b141cf6f3b2e4ae97d6b77d730ab796fd070172))
* **server:** pastebin creation now adds +1 to the user created stats ([3170bba](https://github.com/shoshta73/homehub/commit/3170bbab073941f2875469ce4e5bf7e0ba7107b6))
* **server:** pastebin model ([0e26275](https://github.com/shoshta73/homehub/commit/0e26275c25f5c16cd02e6d62836fd5fb670c4c29))
* **server:** pastebin/create handler ([9a9eaca](https://github.com/shoshta73/homehub/commit/9a9eacacb8c4e587324a88d8845a70dae730382b))
* **server:** serve frontend ([11382f7](https://github.com/shoshta73/homehub/commit/11382f7e1db926fce93e807f49a18b999ddd5095))
* **server:** user model ([f3cd281](https://github.com/shoshta73/homehub/commit/f3cd28129810a9008b3cecfc7ba0492b6f0b57a4))
* **server:** user permissions ([95d514f](https://github.com/shoshta73/homehub/commit/95d514fdf876b3957f8d10e8f9f485a533709ae1))

### Documentation

* setup basic structure ([4a95efe](https://github.com/shoshta73/homehub/commit/4a95efeeaf91b9ef47566c3428fdedf571a96260))

### Miscellaneous Chores

* add dockerignore file ([9edc5c2](https://github.com/shoshta73/homehub/commit/9edc5c26e352acc573593b5971fd233bfdfdfe42))
* add frontend dockerfile ([fe8e832](https://github.com/shoshta73/homehub/commit/fe8e8328d1fa81e234ed5926d73b03ea6197d227))
* add nginx config ([73900dd](https://github.com/shoshta73/homehub/commit/73900dd38fe3c052c7bae77c9a69ad851c51c175))
* add server dockerfile ([ee03175](https://github.com/shoshta73/homehub/commit/ee031751e88f9c12067c89edaff868f371700d95))
* create clean script ([4263f85](https://github.com/shoshta73/homehub/commit/4263f858a49483788a0c1cfa9febcff5e1f54e21))
* create gitignore ([b4d39e4](https://github.com/shoshta73/homehub/commit/b4d39e4a16472cf1eed52d7d12551b1b496ba0a1))
* **docs:** md book init ([e57c167](https://github.com/shoshta73/homehub/commit/e57c167acde8f0b0e7710a53544470b3e6f7045e))
* **frontend:** add rollup build visualizer ([541b010](https://github.com/shoshta73/homehub/commit/541b010abe44fb9412d87adf321b4701322730ab))
* **frontend:** add zustand ([d8d19d5](https://github.com/shoshta73/homehub/commit/d8d19d50dd64547b310b0e0a61c277c2f017a34c))
* **frontend:** added menubar ([6f4006f](https://github.com/shoshta73/homehub/commit/6f4006f865b1b7d7629779b417d161b8214c9141))
* **frontend:** added shacn textarea ([993fd0a](https://github.com/shoshta73/homehub/commit/993fd0ab4a66aecee736a73d992d9fba4408059a))
* **frontend:** added shadcn form ([1930dca](https://github.com/shoshta73/homehub/commit/1930dca4325cca8dcc785b7372f52c018e5b4024))
* **frontend:** added shadcn input ([8e4957c](https://github.com/shoshta73/homehub/commit/8e4957c170b2d291d02a66961b4f6de343814d7a))
* **frontend:** fix formatting ([cad3416](https://github.com/shoshta73/homehub/commit/cad3416ab82dc8c438c5e3a14d34c415b46c3050))
* **frontend:** fix typo in HomeView ([488cf10](https://github.com/shoshta73/homehub/commit/488cf10566c2ab8ff36e5c462a64b2bd34717702))
* **frontend:** remove boilerplate ([8b67c01](https://github.com/shoshta73/homehub/commit/8b67c01925cb105e679f44fac59b2bdec794baab))
* **frontend:** setup react router ([b470ecb](https://github.com/shoshta73/homehub/commit/b470ecb52c6c7c8902dda2635a7f2dd768147ff8))
* **frontend:** setup vitest ([fd2cb4f](https://github.com/shoshta73/homehub/commit/fd2cb4f843687c3887500ea23e9bfb4cc22d7d72))
* initialize frontend project ([b89ee4b](https://github.com/shoshta73/homehub/commit/b89ee4b7f6890c13948e28a9431a153f74c3845a))
* remove boilerplatecode ([b17e5ed](https://github.com/shoshta73/homehub/commit/b17e5edbeeb920d3c15ab988a3d71cb22e1cf71a))
* **server:** cleanup go.mod ([7817a2d](https://github.com/shoshta73/homehub/commit/7817a2ddc67ecf50a354592b61d1c9f0d7a94c82))
* setup echo server ([a5097bf](https://github.com/shoshta73/homehub/commit/a5097bf844c0dbf0b6303f3bcb7a4429254ef0d2))
* setup go project ([c7a89bf](https://github.com/shoshta73/homehub/commit/c7a89bfc72041db0e6b4946eae5d148ea3884cff))
* setup release it ([3961d3c](https://github.com/shoshta73/homehub/commit/3961d3cae838873b211eba80279dadc7d03486c2))
* setup shadcn ([b23b575](https://github.com/shoshta73/homehub/commit/b23b575040be320daef34bc5da81bdfd6d3ca8b6))
* setup tailwindcss ([b3b5cad](https://github.com/shoshta73/homehub/commit/b3b5cad9bd43d5c9fcaa678c0a7c7d7da7c33c26))
* update .editorconfig ([999360b](https://github.com/shoshta73/homehub/commit/999360b3e52306ce7d3fca1050fdb4c6b7125ee3))
* update release-it config ([7aee57d](https://github.com/shoshta73/homehub/commit/7aee57d8e1cda2972079d3db0b52fae8b73ee284))
* update release-it configuration ([208b2b4](https://github.com/shoshta73/homehub/commit/208b2b47ffa3b8ce5dc2d7bf7b90fb785f373c6e))

### Code Refactoring

* **frontend:** lazy import pastebinView ([49ff303](https://github.com/shoshta73/homehub/commit/49ff3033664013305a7528c75776134bd0037d11))
* **frontend:** move menubar to its own component ([623ab4e](https://github.com/shoshta73/homehub/commit/623ab4e12feafa72ec078f6c7eb00f02c74f1565))
* **frontend:** use browser router ([a79716a](https://github.com/shoshta73/homehub/commit/a79716a55a4b8080883b0b7760761921fda5fc1f))
* **server:** cleanup identicon file ([9259206](https://github.com/shoshta73/homehub/commit/92592066fa34dff04e93964bc1bdd4e9841afc56))
* **server:** move data dir constant to constants module ([a30e304](https://github.com/shoshta73/homehub/commit/a30e304b4acbf476bf4b6160a8bf6a3a88b6e27f))

### Build System

* **frontend:** improve manuall chunking ([1749f0c](https://github.com/shoshta73/homehub/commit/1749f0c0d9aa04043302814e6d3378c44770fc5f))
* **frontend:** manually chuck react-router ([4af9d44](https://github.com/shoshta73/homehub/commit/4af9d442f1e6dfd812a71a76873f3b14272ea815))
* **frontend:** manually chunk shadcn ([ab5f7b6](https://github.com/shoshta73/homehub/commit/ab5f7b68d919a0c4582dd4a5074534b25cac8619))
* **frontend:** put views in to their dis directory ([c8d649b](https://github.com/shoshta73/homehub/commit/c8d649b8094aa82ee7c4aa1a19e0a6d05ce167a2))
* **frontend:** update dockerfile ([3029098](https://github.com/shoshta73/homehub/commit/302909800e6b7ed8a582a74c90ab9d6c7cfd6d3d))
* **frontend:** update dockerfile ([c045e7f](https://github.com/shoshta73/homehub/commit/c045e7fd56bdf9bb64ea16965dc855b6eb2b06f8))
* **frontend:** update manual chunks ([47ffd68](https://github.com/shoshta73/homehub/commit/47ffd683a706783ad57dcda950a1aa746d3c6b4c))
* move to single docker files ([3db6735](https://github.com/shoshta73/homehub/commit/3db6735145598603dab5212d56e0241f88440f18))

### CI - CD

* add dockerhub ([4a3db2a](https://github.com/shoshta73/homehub/commit/4a3db2ad424559504917961ce33e2b76ef452e24))
* create build jobs ([b0413c3](https://github.com/shoshta73/homehub/commit/b0413c3c73483c0e7ce4ca99ca180ce0300672dd))
* create release workflow ([84387ea](https://github.com/shoshta73/homehub/commit/84387eacab399ca9567debb695a39973258492b1))
* fix workflow permissions ([b5e9fa9](https://github.com/shoshta73/homehub/commit/b5e9fa9d5031b27e62f9649d84dbd1c8651c1fb3))
* use caching ([6b44411](https://github.com/shoshta73/homehub/commit/6b4441169edf65972e365f5991a10421c6f54a00))

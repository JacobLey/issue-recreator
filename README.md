# Issue - SWC resolveFully missing functionality

1. Open this package in devcontainer
2. Run `pnpm i`
3. Run `./node_modules/.bin/swc ./src -d ./dist --config-file .swcrc --strip-leading-paths -C module.type=commonjs -C module.ignoreDynamic=true --only '**/*.cts' --out-file-extension cjs`
  * This builds _only_ the CJS file
4. Run `./node_modules/.bin/swc  ./src -d ./dist --copy-files --config-file .swcrc --strip-leading-paths --only '**/*.ts'`
5. Inspect `/dist/bar.js`
   * Note that the import path has been updated to point to the ESM file (instead of original CJS file), which will result in logically different code execution
6. Run `./node_modules/.bin/swc  ./src -d ./dist --copy-files --config-file .swcrc-noresolve --strip-leading-paths --only '**/*.ts'`
7. Inspect `/dist/bar.js`
   * Note that the import path has been truncated, which is not allowed in ESM
8. Dream of a third option that doesn't touch resolution paths at all
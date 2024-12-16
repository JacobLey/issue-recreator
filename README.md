# Issue - SWC resolveFully missing functionality

1. Open this package in devcontainer
2. Run `pnpm i`
3. Run `./node_modules/.bin/swc  ./src -d ./dist --copy-files --config-file .swcrc --strip-leading-paths --only '**/*.ts'`
4. Inspect `/dist/bar.js`
   * Note that the import path has been updated to point to the ESM file (instead of original CJS file), which will result in logically different code execution
5. Run `./node_modules/.bin/swc  ./src -d ./dist --copy-files --config-file .swcrc-noresolve --strip-leading-paths --only '**/*.ts'`
6. Inspect `/dist/bar.js`
   * Note that the import path has been truncated, which is not allowed in ESM
7. Dream of a third option that doesn't touch resolution paths at all
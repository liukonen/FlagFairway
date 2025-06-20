const esbuild = require('esbuild')
const { sassPlugin } = require('esbuild-sass-plugin')
const fs = require('fs')
const path = require('path')

const copyStaticPlugin = {
  name: 'copy-static',
  setup(build) {
    build.onEnd(() => {
      const src = path.resolve(__dirname, 'static')
      const dest = path.resolve(__dirname, 'build')
      function copyRecursiveSync(srcDir, destDir) {
        if (!fs.existsSync(srcDir)) return
        const stats = fs.statSync(srcDir)
        if (stats.isDirectory()) {
          if (!fs.existsSync(destDir)) fs.mkdirSync(destDir, { recursive: true })
          for (const file of fs.readdirSync(srcDir)) {
            copyRecursiveSync(path.join(srcDir, file), path.join(destDir, file))
          }
        } else {
          fs.copyFileSync(srcDir, destDir)
        }
      }
      copyRecursiveSync(src, dest)
      console.log('Copied static assets to build.')
    })
  }
}

esbuild.build({
  entryPoints: ['src/index.tsx'],
  bundle: true,
  outfile: 'build/bundle.js',
  minify: true,
  sourcemap: true,
  jsx: 'automatic',
  jsxImportSource: 'preact',
  plugins: [sassPlugin(), copyStaticPlugin],
  define: { 'process.env.NODE_ENV': '"production"' },
  target: ['esnext'],
}).catch(() => process.exit(1))
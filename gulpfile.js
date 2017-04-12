var gulp = require('gulp');
var browserify = require('browserify');
var vinyl = require('vinyl-source-stream');
var tsify = require('tsify');
var watchify = require('watchify');
var gutil = require('gulp-util');

const JS_OUTPUT = 'bundle.js';
const DIST_DIR = './dist';

function compile() {
  return browserify({
    basedir: '.',
    debug: true,
    entries: ['src/main.ts'],
    cache: {},
    packageCache: {},
  }).plugin(tsify);
}

// Compile TypeScript and bundle everything in a single bundle.js file.
function build() {
  return compile()
    .bundle()
    .pipe(vinyl(JS_OUTPUT))
    .pipe(gulp.dest(DIST_DIR));
}

// Watch the source code and recompile on modifications
var watchedBrowserify = watchify(compile());
watchedBrowserify.on("update", watch);
watchedBrowserify.on("log", gutil.log);

function watch() {
  return watchedBrowserify
    .bundle()
    .pipe(vinyl(JS_OUTPUT))
    .pipe(gulp.dest(DIST_DIR));
}

gulp.task('default', watch);
gulp.task('build', build);

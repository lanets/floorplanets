'use strict'

var gulp        = require('gulp'),
    browserify  = require('browserify'),
    vinyl       = require('vinyl-source-stream'),
    tsify       = require('tsify'),
    watchify    = require('watchify'),
    gutil       = require('gulp-util'),
    es          = require('event-stream'),
    rename      = require('gulp-rename');


// path of the applications entry points
const APPS = [
    './src/main.ts',
];


// Build the project
gulp.task('build', () => {

    // Create a stream for each application entry point
    const streams = APPS.map((app) => {

        gutil.log('Begin build for', app);

        const bundler = browserify({
            debug: true,
            entries: [app],
            cache: {},
            packageCache: {},
        }).plugin(tsify);

        const filename = app.replace(/^.*[\\\/]/, '');
        return wrapBundler(bundler, filename);
    });

    return es.merge(streams);
});

// Watch the source code for modifications, recompile the application on change.
gulp.task('default', () => {

    const streams = APPS.map((app) => {

        let bundler = browserify({
            debug: true,
            entries: [app],
            cache: {},
            packageCache: {},
        }).plugin(tsify);

        bundler = watchify(bundler);

        const filename = app.replace(/^.*[\\\/]/, '');
        const watchfn = getBundlerHandler(bundler, filename);
        bundler.on('update', watchfn);
        bundler.on('log', gutil.log);

        return watchfn(); // Compile the bundler
    });

    return es.merge(streams);
});

function wrapBundler(bundler, filename) {
    return bundler
        .bundle()
        .pipe(vinyl(filename))
        .pipe(rename({ extname: '.bundle.js' }))
        .pipe(gulp.dest('./dist'));
}

function getBundlerHandler(bundler, filename) {
  return () => {
    gutil.log('Begin build for', filename);
    return wrapBundler(bundler, filename);
  }
}

module.exports = function(grunt) {

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    bower: {
      install: {
        options: {
          targetDir: './assets/libs'
        }
      }
    },

    jshint: {
      all: [ 'Gruntfile.js', 'app/*.js', 'app/**/*.js' ]
    },

    html2js: {
      dist: {
        src: [ 'app/components/**/*.html' ],
        dest: 'tmp/templates.js'
      },
      options: {
        mode: 'compress',
        wrap: false
      }
    },

    concat: {
      options: {
        separator: ';'
      },
      dist: {
        src: [ 'app/app.module.js', 'app/*.js', 'app/components/**/*.js', 'tmp/*.js' ],
        dest: 'dist/app.js'
      }
    },

    uglify: {
      dist: {
        files: {
          'dist/app.js': [ 'dist/app.js' ]
        },
        options: {
          mangle: true
        }
      }
    },

    clean: {
      temp: {
        src: [ 'tmp' ]
      }
    },

    watch: {
      dev: {
        files: [ 'Gruntfile.js', 'app/*.js', 'app/**/*.js', 'app/**/*.html' ],
        tasks: [ 'jshint', 'html2js:dist', 'concat:dist', 'clean:temp' ],
        options: {
          atBegin: true
        }
      }
    }
  });

  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-html2js');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-bower-task');

  grunt.registerTask('dev', ['bower', 'watch:dev']);
  grunt.registerTask('package', ['bower', 'jshint', 'html2js:dist', 'concat:dist', 'uglify:dist','clean:temp']);
};

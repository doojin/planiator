module.exports = function(grunt) {

    var tasks = [
        'grunt-contrib-sass',
        'grunt-contrib-copy',
        'grunt-contrib-qunit',
        'grunt-bower-task',
        'grunt-contrib-uglify',
        'grunt-contrib-cssmin'
    ];

    tasks.forEach(function(task) {
        grunt.loadNpmTasks(task);
    });

    grunt.initConfig({

        // grunt-contrib-sass
        sass: {
            app: {
                options: {
                    loadPath: [
                        'lib/foundation/scss',
                        'lib/foundation/scss/foundation'
                    ]
                },
                files: {
                    'assets/css/app.css': 'lib/foundation/scss/app.scss'
                }
            }
        },

        // grunt-contrib-copy
        copy: {
            foundation: {
                files: [
                    {expand: true, src: ['scss/**'], dest: 'lib/foundation'}
                ]
            },
            dependencies: {
                files: [
                    // Timepicker
                    {flatten: true, expand: true, src: 'lib/jt.timepicker/*', dest: 'assets/vendor/timepicker'},
                    // Foundation icons
                    {flatten: true, expand: true, src: 'lib/foundation-icon-fonts/svgs/*', dest: 'assets/vendor/icons/svgs'},
                    {flatten: true, expand: true, src: 'lib/foundation-icon-fonts/foundation-icons.*', dest: 'assets/vendor/icons/'},
                    // Foundation
                    {flatten: true, expand: true, src: 'lib/foundation/js/foundation.min.js', dest: 'assets/vendor/foundation/js'},
                    {flatten: true, expand: true, src: 'lib/foundation/js/foundation/foundation.topbar.js', dest: 'assets/vendor/foundation/js'},
                    {flatten: true, expand: true, src: 'lib/foundation/js/foundation.min.js', dest: 'assets/vendor/foundation/js'},
                    // jQuery & jQuery UI
                    {flatten: true, expand: true, src: 'lib/jquery/jquery.js', dest: 'assets/vendor/jquery'},
                    {flatten: true, expand: true, src: 'lib/jquery-ui/jquery-ui.min.js', dest: 'assets/vendor/jquery'},
                    {flatten: true, expand: true, src: 'lib/jquery-ui/ui/minified/datepicker.min.js', dest: 'assets/vendor/jquery'},
                    {flatten: true, expand: true, src: 'lib/jquery-ui/themes/ui-darkness/images/*', dest: 'assets/vendor/jquery/images'},
                    {flatten: true, expand: true, src: 'lib/jquery-ui/themes/ui-darkness/jquery-ui.min.css', dest: 'assets/vendor/jquery'},
                    // BlockUI
                    {flatten: true, expand: true, src: 'lib/blockui/*', dest: 'assets/vendor/blockui'},
                    // ColorPicker
                    {
                        flatten: true,
                        expand: true,
                        src: 'node_modules/evol.colorpicker/css/evol.colorpicker.min.css',
                        dest: 'assets/vendor/colorpicker'
                    },
                    {
                        flatten: true,
                        expand: true,
                        src: 'node_modules/evol.colorpicker/js/evol.colorpicker.min.js',
                        dest: 'assets/vendor/colorpicker'
                    }
                ]
            }
        },

        // grunt-contrib-qunit
        qunit: {
            all: ['tests/*.html']
        },

        // grunt-bower-task
        bower: {
            install: {
                options: {
                    targetDir: 'lib',
                    cleanBowerDir: true,
                    cleanTargetDir: true,
                    layout: "byComponent"
                }
            }
        },

        // grunt-contrib-uglify
        uglify: {
            dependencies: {
                files: {
                    'assets/vendor/foundation/js/foundation.topbar.min.js': ['assets/vendor/foundation/js/foundation.topbar.js'],
                    'assets/vendor/jquery/jquery.min.js': ['assets/vendor/jquery/jquery.js'],
                    'assets/vendor/timepicker/jquery.timepicker.min.js': ['assets/vendor/timepicker/jquery.timepicker.js'],
                    'assets/vendor/blockui/jquery.blockUI.min.js': ['assets/vendor/blockui/jquery.blockUI.js']
                }
            }
        },

        // grunt-contrib-cssmin
        cssmin: {
            dependencies: {
                files: {
                    'assets/vendor/icons/foundation-icons.min.css': ['assets/vendor/icons/foundation-icons.css'],
                    'assets/vendor/timepicker/jquery.timepicker.min.css': ['assets/vendor/timepicker/jquery.timepicker.css']
                }
            },
            appstyles: {
                files: {
                    'assets/css/app.min.css': ['assets/css/app.css']
                }
            }
        }
    });

    grunt.registerTask('scss', ['copy:foundation', 'sass', 'cssmin:appstyles']);
    grunt.registerTask('test', ['qunit']);
    grunt.registerTask('deps', ['bower:install', 'copy:dependencies', 'uglify:dependencies', 'cssmin:dependencies']);
};

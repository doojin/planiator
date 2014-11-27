module.exports = function(grunt) {
    
    var tasks = [
        'grunt-contrib-sass'
    ];
    
    tasks.forEach(function(task) {
        grunt.loadNpmTasks(task);
    });
    
    grunt.initConfig({
        sass: {
            foundation: {
                options: {
                    loadPath: [
                        'libs/foundation/scss',
                    ]
                },
                files: {
                    'assets/css/foundation.css': 'libs/foundation/scss/foundation.scss'
                }
            },
            customSCSSFiles: {
                options: {
                    loadPath: [
                        'libs/foundation/scss',
                    ]
                },
                files: {
                    'assets/css/common.css': 'scss/common.scss',
                    'assets/css/index.css': 'scss/index.scss',
                    'assets/css/calendar.css': 'scss/calendar.scss'
                }
            }
        }
    });
    
    grunt.registerTask('scss', ['sass']);
    grunt.registerTask('scss_custom', ['sass:customSCSSFiles']);
    
};
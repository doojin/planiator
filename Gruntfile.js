module.exports = function(grunt) {

    var tasks = [
        'grunt-contrib-sass',
        'grunt-contrib-copy'
    ];

    tasks.forEach(function(task) {
        grunt.loadNpmTasks(task);
    });

    grunt.initConfig({
        sass: {
            app: {
                options: {
                    loadPath: [
                        'libs/foundation/scss',
                        'libs/foundation/scss/foundation'
                    ]
                },
                files: {
                    'assets/css/app.css': 'libs/foundation/scss/app.scss'
                }
            }
        },
        copy: {
            foundation: {
                files: [
                    {expand: true, src: ['scss/**'], dest: 'libs/foundation'}    
                ]
            }
        }
    });

    grunt.registerTask('scss', ['copy', 'sass']);

};

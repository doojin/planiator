module.exports = function(grunt) {
    
    var tasks = [
        'grunt-contrib-sass'
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
        }
    });
    
    grunt.registerTask('scss', ['sass']);
    
};
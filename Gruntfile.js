module.exports = function(grunt){

	var jsFilePaths = [
		'public/js/*.js'
	];

	grunt.initConfig({
		pkg: grunt.file.readJSON('package.json'),
		appDir: 'public',
		builtDir: 'public/dist',
		requirejs:{
			main:{
				options:{
					appDir: '<%= appDir %>',
					baseUrl: './js',
					dir: '<%= builtDir %>',
					optimizeCss: 'none', // will be taken care of with compass
					optimize: 'none' // will be taken care of with an uglify task
				}
			}
		},
		uglify:{
			options:{
				banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> - <%= pkg.author %> */\n'
			},
			build:{
				files:(function(){
					var files = [];
					jsFilePaths.forEach(function(val){
						files.push({
							expand: true,
							cwd: '<%= builtDir %>',
							src: val,
							dest: '<%= builtDir %>'
						});
					});

					return false;
				})()
			}
		},
		jshint: {
			options:{
				reporter: require('jshint-stylish')
			},
			all:[
				'Gruntfile.js',
				'<%= appDir %>/js/{,*/}*.js'
			]
		},
		compass:{
			dist:{
				options: {
					sassDir: '<%= appDir %>/sass',
					cssDir: '<%= appDir %>/css',
					banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> - <%= pkg.author %> */\n',
					specify: '<%= appDir %>/sass/*.scss',
					imagesDir: 'public/imgs',
					fontsDir: 'public/fonts',
					outputStyle:'compressed',
					noLineComments: true,
					environment: 'production'
				}
			},
			dev:{
				options: {
					sassDir: '<%= appDir %>/sass',
					cssDir: '<%= appDir %>/css',
					banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> - <%= pkg.author %> */\n',
					specify: '<%= appDir %>/sass/*.scss',
					imagesDir: 'public/imgs',
					fontsDir: 'public/fonts',
					outputStyle:'expanded',
					noLineComments: false,
					environment: 'development'
				}
			}
		},
		watch:{
			scripts:{
				files: (function(){
					var files = [];
					jsFilePaths.forEach(function(val){
						files.push('<%= appDir %>/'+val);
					});
					return files;
				})(),
				tasks:['jshint'],
				options:{
					spawn: false
				}
			},
			compass:{
				files: '<%= appDir %>/sass/*.scss',
				tasks: ['compass:dev'],
				options:{
					spawn:false
				}
			}
		},
		go:{
			myapp:{
				output:'dynamic',
				cmd:'goapp',
				build_flags: ['-tags','appengine'],
				run_files: ['index.go']
			}
		},
		concurrent:{
			prod:{
				tasks:['go:run:myapp','watch:scripts','watch:compass']
			},
			options:{
				logConcurrentOutput: true
			}
		}
	});

	grunt.loadNpmTasks('grunt-contrib-requirejs');
	grunt.loadNpmTasks('grunt-contrib-jshint');
	grunt.loadNpmTasks('grunt-contrib-jasmine');
	grunt.loadNpmTasks('grunt-contrib-concat');
	grunt.loadNpmTasks('grunt-contrib-uglify');
	grunt.loadNpmTasks('grunt-contrib-watch');
	grunt.loadNpmTasks('grunt-contrib-compass');
	grunt.loadNpmTasks('grunt-shell');
	grunt.loadNpmTasks('grunt-concurrent');
	grunt.loadNpmTasks('grunt-go');

	grunt.registerTask('default',['jshint','compass:dev','concurrent:prod']);

	grunt.registerTask('prod', ['jshint','requirejs','uglify','compass:dist']);
};

if(typeof define !== 'function'){
	var define = require('amdefine')(module);
}

require.config({
	paths: {
		'jquery': '../vendor/jquery/dist/jquery.min',
		'respondJS': '../vendor/respondJS/dest/respond.min',
		'bootstrap': '../vendor/bootstrap/dist/js/bootstrap.min',
		'nprogress':'../vendor/nprogress/nprogress',
		'holder':'../vendor/holderjs/holder',
		'alertify':'../vendor/alertify.js/lib/alertify.min',
		'banner':'./banner',
		'content':'./content',
		'quote':'./quote',
		'templates':'../views'
	},

	shim: {
		'jquery':{
			'exports':'$'
		},
		'bootstrap':['jquery']
	},
	waitSeconds: 15,
	urlArgs: 'bust=v0.1.1'
});

require([
	'require',
	'jquery',
	'alertify',
	'bootstrap',
	'respondJS',
	'nprogress',
	'holder'], function(require, $, alertify){

		var querySearch = function (name) {
			return unescape(window.location.search.replace(new RegExp("^(?:.*[&\\?]" + escape(name).replace(/[\.\+\*]/g, "\\$&") + "(?:\\=([^&]*))?)?.*$", "i"), "$1"));
		};

		var path = window.location.pathname.split('/');
		if(path.length > 2 && path[2] !== ""){
			var page = path[2];
			switch(page){
				case 'banners':
					require(['banner']);
					break;
				case 'content':
					require(['content']);
					break;
				case 'quotes':
					require(['quote']);
					break;
				default:
					break;
			}
		}

		/* Global URL Prompting */
		var err = querySearch('error');
		var success = querySearch('success');
		if(err.length > 0){
			alertify.error(err);
		}
		if(success.length > 0){
			alertify.success(success);
		}

});

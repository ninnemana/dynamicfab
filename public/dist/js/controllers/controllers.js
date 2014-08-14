define([
	'jquery',
	'nprogress',
	'routes/routes',
	'controllers/AppController',
	'controllers/HomeController',
	'controllers/AboutController',
	'controllers/EquipmentController',
	'controllers/WarrantyController',
	'controllers/PolicyController',
	'controllers/TermsController',
	'controllers/QuoteController',
	'controllers/TestimonialsController'],
	function($, NProgress, routes, app, home, about, equipment, warranty, policy, terms, quote, testimonials){

		var controllers = {
			home: home,
			warranty: warranty,
			about: about,
			equipment: equipment,
			policy: policy,
			terms: terms,
			quote: quote,
			testimonials: testimonials
		};

		var setUpRoutes = function(angModule){
			angModule.config(function($routeProvider, $locationProvider, $interpolateProvider) {
				$.each(routes, function(key, val) {
					$routeProvider.when(val.route,{
						template: val.template,
						controller: val.controller,
						title: val.title
					});
					$routeProvider.otherwise({ redirectTo: "/home" });
				});

				$interpolateProvider.startSymbol('[[').endSymbol(']]');
				$locationProvider.html5Mode(true);
				$routeProvider.otherwise({ redirectTo: routes.home.route });
			});
			angModule.run(function($rootScope, $location){
				$rootScope.$on('$routeChangeStart',function(){
					NProgress.start();
				});
				$rootScope.$on('$routeChangeSuccess', function(next, last){
					NProgress.done();
				});

				$rootScope.isActive = function(viewPath){
					return viewPath == $location.path();
				};
			});
		};

		var initialize = function(angModule){
			angModule.controller('AppController', app);
			$.each(controllers, function(name, ctrl) {
				angModule.controller(name, ctrl);
			});
			setUpRoutes(angModule);
		};

		return {
			initialize: initialize
		};
});

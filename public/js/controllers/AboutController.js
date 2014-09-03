define(['/js/vendor/angular-sanitize/angular-sanitize.min.js'],function(){
	'use strict';
	var ctlr = ['$scope','$sce', '$http', function($scope, $sce, $http){
		$scope.about = {};
		$http({
			method:'get',
			url:'/api/aboutus'
		}).success(function(data,status){
			$scope.about = data;

		});
		$scope.trustAboutContent = function() {
			return $sce.trustAsHtml($scope.about.body);
		};
	}];

	return ctlr;
});
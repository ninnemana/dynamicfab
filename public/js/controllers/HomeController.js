define([],function(){
	'use strict';
	var ctlr = ['$scope','BannerService', 'ContentService', function($scope, bs, cs){

		$scope.cols = [];
		cs.query(function(content){
			$scope.cols = content;
		});

		$scope.carousel = [];
		bs.query(function(banners){
			$scope.carousel = banners;
		});
	}];

	return ctlr;
});
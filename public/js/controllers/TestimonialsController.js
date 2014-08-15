define([],function(){
	'use strict';
	var ctlr = ['$scope', 'TestimonialService', function($scope, ts){
		$scope.testimonials = [];
		ts.query(function(ts){
			$scope.testimonials = ts;
		});
	}];

	return ctlr;
});
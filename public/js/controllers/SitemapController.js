define([],function(){
	'use strict';
	var ctlr = ['$scope', function($scope, bs, cs){
		var org = window.location.origin;
		$scope.pages = [{
			name:'Home',
			link: org+'/',
		},
		{
			name:'About Us',
			link: org+'/about',
		},{
			name:'Capabilities & Equipment',
			link: org+'/equipment',
		},
		{
			name:'Request a Quote',
			link: org+'/quote',
		},{
			name:'Past Jobs & Testimonials',
			link: org+'/testimonials',
		}];
	}];

	return ctlr;
});
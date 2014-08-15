define([], function () {
	'use strict';

	var service = ['$resource', function ($resource) {
		return $resource('/api/testimonials/:id', {id: '@id'},{
			query:{method:'GET', isArray: true},
			get:{method: 'GET', isArray:false},
			post:{method: 'POST'}
		});
	}];

	return service;
});
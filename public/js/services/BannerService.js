define([], function () {
	'use strict';

	var service = ['$resource', function ($resource) {
		return $resource('/api/banners/:id', {id: '@id'},{
			query:{method:'GET', isArray: true},
			get:{method: 'GET', isArray:false},
			post:{method: 'POST'}
		});
	}];

	return service;
});
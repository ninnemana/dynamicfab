define(['jquery',
	'services/BannerService',
	'services/ContentService',
	'services/TestimonialService'],function($, bs, cs, ts){
	'use strict';

	var services = {
		BannerService: bs,
		ContentService: cs,
		TestimonialService: ts
	};

	var initialize = function(angModule){
		$.each(services, function(name, service) {
			angModule.factory(name, service);
		});
	};

	return {
		initialize: initialize
	};

});
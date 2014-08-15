define(['jquery',
	'services/BannerService',
	'services/ContentService',
	'services/TestimonialService',
	'services/EquipmentService'],function($, bs, cs, ts, es){
	'use strict';

	var services = {
		BannerService: bs,
		ContentService: cs,
		TestimonialService: ts,
		EquipmentService: es
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
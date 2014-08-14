define(['jquery',
	'services/BannerService',
	'services/ContentService'],function($, bs, cs){
	

	var services = {
		BannerService: bs,
		ContentService: cs
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
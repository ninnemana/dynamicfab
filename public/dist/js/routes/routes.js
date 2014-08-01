define([
	'/js/vendor/requirejs-text/text.js!/js/views/home.html',
	'/js/vendor/requirejs-text/text.js!/js/views/about.html',
	'/js/vendor/requirejs-text/text.js!/js/views/equipment.html'
],function(homeTemplate, aboutTemplate, equipmentTemplate){
	return {
		home: {
			title: 'Home',
			route: '/home',
			controller: 'home',
			template: homeTemplate
		},
		about: {
			title: 'About Us',
			route: '/about',
			controller: 'about',
			template: aboutTemplate
		},
		equipment: {
			title: 'Capabilities & Equipment',
			route: '/equipment',
			controller: 'equipment',
			template: equipmentTemplate
		}
	};
});
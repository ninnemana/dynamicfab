define([
	'/js/vendor/requirejs-text/text.js!/js/views/home.html',
	'/js/vendor/requirejs-text/text.js!/js/views/about.html',
	'/js/vendor/requirejs-text/text.js!/js/views/equipment.html',
	'/js/vendor/requirejs-text/text.js!/js/views/quote.html',
	'/js/vendor/requirejs-text/text.js!/js/views/testimonials.html'
],function(homeTemplate, aboutTemplate, equipmentTemplate, quoteTemplate, testimonialsTemplate){
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
		},
		quote: {
			title: 'Request a Quote',
			route: '/quote',
			controller: 'quote',
			template: quoteTemplate
		},
		testimonials: {
			title: 'Past Jobs & Testimonials',
			route: '/testimonials',
			controller: 'testimonials',
			template: testimonialsTemplate
		}
	};
});
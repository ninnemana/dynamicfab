define([],function(){
	'use strict';
	var ctlr = ['$scope', '$http', '$anchorScroll', function($scope, $http, $anchorScroll){
		$scope.heading = "";
		$scope.nameError = null;
		$scope.emailError = null;
		$scope.phoneError = null;
		$scope.descError = null;
		$scope.formError = null;
		$scope.formSuccess = null;
		$scope.quote = {
			name:'',
			email:'',
			phone:'',
			desc:'',
			created: null,
			validate: function(){
				var err = false;
				$scope.nameError = null;
				$scope.emailError = null;
				$scope.phoneError = null;
				$scope.descError = null;
				$scope.formError = null;

				if($scope.quote === undefined){
					$scope.formError = 'Failed to submit the form, please try back.';
					err = true;
					return;
				}
				if($scope.quote.desc === undefined || $scope.quote.desc === ''){
					$('#desc').focus();
					$scope.descError = 'Please describe the project for which you are requesting a quote.';
					err = true;
				}
				if(($scope.quote.email === undefined || $scope.quote.email === '') && ($scope.quote.phone === undefined || $scope.quote.phone === '')){
					$('#email').focus();
					$scope.emailError = 'Email or phone number is required';
					$scope.phoneError = $scope.emailError;
					err = true;
				}
				if($scope.quote.name === undefined || $scope.quote.name === ''){
					$('#name').focus();
					$scope.nameError = 'Name is required';
					err = true;
				}

				if(err){
					$scope.formError = 'Failed to request quote.';
				}

				return err;
			}
		};
		$scope.requestQuote = function(e){
			if($scope.quote.validate()){
				return false;
			}

			$http({
				method:'POST',
				url: '/api/quote',
				data: $scope.quote
			}).success(function(data, status) {
				$scope.quote.name = '';
				$scope.quote.company = '';
				$scope.quote.address = '';
				$scope.quote.city = '';
				$scope.quote.state = '';
				$scope.quote.zip = '';
				$scope.quote.phone = '';
				$scope.quote.fax = '';
				$scope.quote.email = '';
				$scope.quote.website = '';
				$scope.quote.request_date = '';
				$scope.quote.delivery_date = '';
				$scope.quote.part = '';
				$scope.quote.quantity = '';
				$scope.quote.cost = '';
				$scope.quote.pkg = '';
				$scope.quote.desc = '';
				$scope.quote.created = null;
				$scope.formSuccess = true;
				$anchorScroll();
			}).error(function(data, status) {
				$scope.formError = data;
			});
		};

		$scope.countries = [{"country_id":1,"country":"United States","abbreviation":"US","states":[{"state_id":1,"state":"Alabama","abbreviation":"AL"},{"state_id":2,"state":"Alaska","abbreviation":"AK"},{"state_id":3,"state":"American Samoa","abbreviation":"AS"},{"state_id":4,"state":"Arizona","abbreviation":"AZ"},{"state_id":5,"state":"Arkansas","abbreviation":"AR"},{"state_id":6,"state":"California","abbreviation":"CA"},{"state_id":7,"state":"Colorado","abbreviation":"CO"},{"state_id":8,"state":"Connecticut","abbreviation":"CT"},{"state_id":9,"state":"Delaware","abbreviation":"DE"},{"state_id":10,"state":"District Of Columbia","abbreviation":"DC"},{"state_id":11,"state":"Federated States Of Micronesia","abbreviation":"FM"},{"state_id":12,"state":"Florida","abbreviation":"FL"},{"state_id":13,"state":"Georgia","abbreviation":"GA"},{"state_id":14,"state":"Guam","abbreviation":"GU"},{"state_id":15,"state":"Hawaii","abbreviation":"HI"},{"state_id":16,"state":"Idaho","abbreviation":"ID"},{"state_id":17,"state":"Illinois","abbreviation":"IL"},{"state_id":18,"state":"Indiana","abbreviation":"IN"},{"state_id":19,"state":"Iowa","abbreviation":"IA"},{"state_id":20,"state":"Kansas","abbreviation":"KS"},{"state_id":21,"state":"Kentucky","abbreviation":"KY"},{"state_id":22,"state":"Louisiana","abbreviation":"LA"},{"state_id":23,"state":"Maine","abbreviation":"ME"},{"state_id":24,"state":"Marshall Islands","abbreviation":"MH"},{"state_id":25,"state":"Maryland","abbreviation":"MD"},{"state_id":26,"state":"Massachusetts","abbreviation":"MA"},{"state_id":27,"state":"Michigan","abbreviation":"MI"},{"state_id":28,"state":"Minnesota","abbreviation":"MN"},{"state_id":29,"state":"Mississippi","abbreviation":"MS"},{"state_id":30,"state":"Missouri","abbreviation":"MO"},{"state_id":31,"state":"Montana","abbreviation":"MT"},{"state_id":32,"state":"Nebraska","abbreviation":"NE"},{"state_id":33,"state":"Nevada","abbreviation":"NV"},{"state_id":34,"state":"New Hampshire","abbreviation":"NH"},{"state_id":35,"state":"New Jersey","abbreviation":"NJ"},{"state_id":36,"state":"New Mexico","abbreviation":"NM"},{"state_id":37,"state":"New York","abbreviation":"NY"},{"state_id":38,"state":"North Carolina","abbreviation":"NC"},{"state_id":39,"state":"North Dakota","abbreviation":"ND"},{"state_id":40,"state":"Northern Mariana Islands","abbreviation":"MP"},{"state_id":41,"state":"Ohio","abbreviation":"OH"},{"state_id":59,"state":"Oklahoma","abbreviation":"OK"},{"state_id":60,"state":"Oregon","abbreviation":"OR"},{"state_id":44,"state":"Palau","abbreviation":"PW"},{"state_id":42,"state":"Pennsylvania","abbreviation":"PA"},{"state_id":43,"state":"Puerto Rico","abbreviation":"PR"},{"state_id":45,"state":"Rhode Island","abbreviation":"RI"},{"state_id":46,"state":"South Carolina","abbreviation":"SC"},{"state_id":47,"state":"South Dakota","abbreviation":"SD"},{"state_id":48,"state":"Tennessee","abbreviation":"TN"},{"state_id":49,"state":"Texas","abbreviation":"TX"},{"state_id":50,"state":"Utah","abbreviation":"UT"},{"state_id":51,"state":"Vermont","abbreviation":"VT"},{"state_id":52,"state":"Virgin Islands","abbreviation":"VI"},{"state_id":53,"state":"Virginia","abbreviation":"VA"},{"state_id":54,"state":"Washington","abbreviation":"WA"},{"state_id":55,"state":"West Virginia","abbreviation":"WV"},{"state_id":56,"state":"Wisconsin","abbreviation":"WI"},{"state_id":57,"state":"Wyoming","abbreviation":"WY"}]},{"country_id":2,"country":"Canada","abbreviation":"CA","states":[{"state_id":61,"state":"Alberta","abbreviation":"AB"},{"state_id":62,"state":"British Columbia","abbreviation":"BC"},{"state_id":63,"state":"Manitoba","abbreviation":"MB"},{"state_id":64,"state":"New Brunswick","abbreviation":"NB"},{"state_id":65,"state":"Newfoundland and Labrador","abbreviation":"NL"},{"state_id":66,"state":"Northwest Territories","abbreviation":"NT"},{"state_id":67,"state":"Nova Scotia","abbreviation":"NS"},{"state_id":68,"state":"Nunavut","abbreviation":"NU"},{"state_id":69,"state":"Ontario","abbreviation":"ON"},{"state_id":70,"state":"Prince Edward Island","abbreviation":"PE"},{"state_id":71,"state":"Quebec","abbreviation":"QC"},{"state_id":72,"state":"Saskatchewan","abbreviation":"SK"},{"state_id":73,"state":"Yukon","abbreviation":"YT"}]},{"country_id":3,"country":"Dominican Republic","abbreviation":"DO","states":[{"state_id":75,"state":"Distrito Nacional","abbreviation":"DN"}]},{"country_id":4,"country":"Mexico","abbreviation":"MX","states":[{"state_id":79,"state":"Distrito Federal","abbreviation":"DF"},{"state_id":76,"state":"Jalisco","abbreviation":"JA"}]},{"country_id":5,"country":"Japan","abbreviation":"JP","states":[{"state_id":77,"state":"Chiba-Ken","abbreviation":"CK"}]},{"country_id":6,"country":"Honduras","abbreviation":"HN","states":[{"state_id":78,"state":"Francisco Morazan","abbreviation":"FMO"}]},{"country_id":7,"country":"South Korea","abbreviation":"KR","states":[{"state_id":80,"state":"Gyeonggi","abbreviation":"GYI"}]},{"country_id":8,"country":"Suriname","abbreviation":"SR","states":[{"state_id":81,"state":"Paramaribo District","abbreviation":"PMD"}]},{"country_id":9,"country":"Egypt","abbreviation":"EG","states":[{"state_id":82,"state":"Cairo Governorate","abbreviation":"CAG"}]},{"country_id":10,"country":"Colombia","abbreviation":"CO","states":[{"state_id":84,"state":"Atlantico","abbreviation":"ATO"},{"state_id":74,"state":"Cundinamarca","abbreviation":"CUN"}]},{"country_id":11,"country":"Taiwan","abbreviation":"TW","states":[{"state_id":83,"state":"Chang-hua","abbreviation":"CH"}]}];

		$http({
			method:'get',
			url:'/api/quote/heading'
		}).success(function(data,status){
			$scope.heading = data;
		});
	}];

	return ctlr;
});
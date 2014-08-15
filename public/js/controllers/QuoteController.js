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
				$scope.quote.phone = '';
				$scope.quote.email = '';
				$scope.quote.desc = '';
				$scope.quote.created = null;
				$scope.formSuccess = true;
				$anchorScroll();
			}).error(function(data, status) {
				$scope.formError = data;
			});
		};

		$http({
			method:'get',
			url:'/api/quote/heading'
		}).success(function(data,status){
			$scope.heading = data;
		});
	}];

	return ctlr;
});
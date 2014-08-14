define([],function(){
	
	var ctlr = ['$scope', function($scope){
		$scope.heading = "All right, Plan B. You might want to watch out the front window there, Larry. Hardly Dude, a new 'vette? The kid's still got, oh, 96 to 97 thousand, depending on the options. These men are nihilists, Donny, nothing to be afraid of. Is this yours, Larry? Is this your homework, Larry? I just want to say, sir, that we're both enormous—on a personal level. Hello, Pilar? My name is Walter Sobchak, we spoke on the phone, this is my associate Jeffrey Lebowski.";
		$scope.nameError = null;
		$scope.emailError = null;
		$scope.phoneError = null;
		$scope.descError = null;
		$scope.formError = null;
		$scope.quote = {
			name:'',
			email:'',
			phone:'',
			desc:'',
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
			if(!$scope.quote.validate()){
				return;
			}
		};
	}];

	return ctlr;
});
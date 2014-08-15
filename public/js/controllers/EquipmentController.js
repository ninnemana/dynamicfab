define([],function(){
	'use strict';
	var ctlr = ['$scope', 'EquipmentService', function($scope, es){
		$scope.groups = [];
		es.query(function(equip){
			$scope.groups = equip;
		});
	}];

	return ctlr;
});
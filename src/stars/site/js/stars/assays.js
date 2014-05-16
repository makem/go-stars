$stars.controller('AssaysCtrl',function($scope, $http, $location){
	$http.defaults.headers.post["Content-Type"] = "application/x-www-form-urlencoded";
	$scope.employees = [];
	$scope.getGenderImageUrl = function(e){
		return e.gender == "M"?"img/avatars/male.png":"img/avatars/female.png";
	};
	$scope.visits = {};
	$scope.select = function(e){
		$scope.selected = e;
		RefreshVisit($scope,$http,e);
	};
	$scope.submit = function(){

		$http.post('/assays/update',$scope.envelope)
			.success(function(data){
				if(data.status=='success'){
					RefreshVisit($scope,$http,$scope.envelope);

				}else{
					alert(data.message)
				}
			})
			.error(function(data){
				alert("Error - "+data);
			})
	};

	$scope.getLabelClass = function(status){
		switch(status){
			case 'notify':
			return 'bg-color-greenLight';
			case 'warning':
			return 'bg-color-blueLight';
			case 'dangerous':
			return 'bg-color-orange';
			case 'critical':
			return 'bg-color-red';
			case 'inactive':
			return 'bg-color-darken';
		}
		return '';
	};
	$scope.getAssayClass=function(status){
		if(status == 'inactive'){
			return 'inactive'
		}
		return "";
	};
	RefreshVisits($scope,$http);

});

function RefreshVisits($scope,$http){
		$http.get('/employees/list')
		.success(function(e){
			$scope.employees = e;
		})
		.error	(function(e){

		});

}

function RefreshVisit($scope,$http,e){
	$http.get('/assays/list/'+e.login)
		.success(function(data){
			$scope.envelope = data;
			$scope.visits = $scope.envelope.visits;
		})
		.error(function(data){
		})
}

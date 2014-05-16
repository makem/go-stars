$stars.controller('EmployeesCtrl',function($scope, $http, $location){
	$http.defaults.headers.post["Content-Type"] = "application/x-www-form-urlencoded";
	$scope.employees = []; 
	$scope.newEmployer = function(){
		$location.path('/employee/new');
	};
	$scope.search='';
	$scope.getGenderImageUrl = function(e){
		return e.gender == "M"?"img/avatars/male.png":"img/avatars/female.png";
	};
	$scope.editEmployee = function(emp){
		$location.path('/employee/'+emp.login);
	};
	$http.get('/employees/list')
		.success(function(e){
			$scope.employees = e;
		})
		.error	(function(e){
			
		});	
});

$stars.controller('EmployeeEditCtrl',function($scope,$http, $location,login){
	$scope.cancel = function(){
		window.history.back();
	};	
	$scope.isNewEmployee = function(){
		return login == "@";
	}; 
	
	$scope.header = $scope.isNewEmployee()?"Регистрация ":"Сотрудник ";
	
	$scope.anyRoleSelected = function(){
		var roles = $scope.employee.roles;
		return Object.keys(roles).some(function(key){
			var value = roles[key].active;
			return value;
		});
	};
	
	$http.get('/employee/data/'+login).success(function(data){
		$scope.employee = data.model;
		$scope.countries = data.countries;
	});

	$scope.submit = function(){
		var url = '/employee/update';
		if($scope.isNewEmployee()){
			url = '/employee/create'
		};
		$http.post(url,$scope.employee)
		.success(function(data){
			if(data.status == "success"){
				$scope.cancel();
			}else{
				alert('Error '+data.message);
			}
			
		})
		.error(function(data){
			alert('Error '+data.message);
		});
	};
	

	$('#birthdate').datepicker({
		dateFormat : 'dd.mm.yy',
		prevText : '<i class="fa fa-chevron-left"></i>',
		nextText : '<i class="fa fa-chevron-right"></i>',
		onSelect : function(selectedDate) {
			$scope.employee.birthDate = selectedDate;
			$scope.$apply();
		}
	});

	$('#registerdate').datepicker({
		dateFormat : 'dd.mm.yy',
		prevText : '<i class="fa fa-chevron-left"></i>',
		nextText : '<i class="fa fa-chevron-right"></i>',
		onSelect : function(selectedDate) {
			$scope.employee.jobStarted = selectedDate;
			$scope.$apply();
		}
	});

});


$stars.controller('EmployeeProfileCtrl',function($scope){
});
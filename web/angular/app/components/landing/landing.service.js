angular.module('landing')
    .service('LandingService', ['$http', function($http) {
            
            this.saveUser = function(user) {
                return $http.post('/user', user);  
            };
            
            this.getUsers = function() {
                return $http.get('/users');
            };
    }]);
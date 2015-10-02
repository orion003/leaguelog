angular.module('home')
    .service('HomeService', ['$http', function($http) {
            
            this.getLeagues = function() {
                return $http.get('/api/leagues');
            };
            
            this.saveUser = function(user) {
                return $http.post('/api/users', user);  
            };
            
    }]);
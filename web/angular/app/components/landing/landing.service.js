angular.module('landing')
    .service('LandingService', ['$http', '$log', function($http, $log) {
            
            this.saveUser = function(user) {
                $log.info("User: " + JSON.stringify(user));
                
                return $http.post('/users', user);  
            };
            
    }]);
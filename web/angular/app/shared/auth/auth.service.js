angular.module('auth')
    .service('AuthService', ['$http', '$localStorage',
        function($http, $localStorage) {

            this.register = function(user, success, error) {
                $http.post('/api/user/register', user).then(
                  function(response) {
                      $localStorage.token = response.data.token;
                      success(response);
                  },
                  error);
            };

            this.login = function(user, success, error) {
              $http.post('/api/user/login', user).then(
                function(response) {
                    $localStorage.token = response.data.token;
                    success(response);
                },
                error);
            };

            this.logout = function(user, success, error) {
                delete $localStorage.token;

                $http.post('/api/user/logout').then(
                  success,
                  error);
            };
    }]);

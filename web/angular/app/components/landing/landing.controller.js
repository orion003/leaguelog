angular.module('landing')
    .controller('LandingController', ['$state', 'LandingService', 'users', 'user',
        function($state, LandingService, users, user){
            this.user = user.data;
            
            this.users = users.data;
            
            this.saveUser = function() {
                LandingService.saveUser(this.user)
                    .then(function() {
                        $state.go('landing') 
                    });
            };
    }]);
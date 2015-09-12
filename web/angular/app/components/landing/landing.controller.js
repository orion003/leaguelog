angular.module('landing')
    .controller('LandingController', ['$state', 'LandingService', 'user',
        function($state, LandingService, user){
            this.user = user.data;

            this.saveUser = function() {
                LandingService.saveUser(this.user)
                    .then(function() {
                        $state.go('landing') 
                    });
            };
    }]);
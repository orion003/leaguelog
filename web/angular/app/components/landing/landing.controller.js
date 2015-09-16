angular.module('landing')
    .controller('LandingController', ['$state', 'LandingService', 'user', '$log',
        function($state, LandingService, user, $log){
            this.user = user.data;

            this.saveUser = function() {
                LandingService.saveUser(this.user).then(
                    function(response) {
                        $state.go('landing.success');
                    }, 
                    function(response) {
                        $log.info(response.data);
                        if(response.data.error == 'user_duplicate_email') {
                            $state.go('landing.duplicate');
                        }
                        else {
                            
                        }
                    });
            };
    }]);
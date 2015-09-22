angular.module('landing')
    .controller('LandingController', ['$state', 'LandingService', '$log',
        function($state, LandingService, $log){
            this.saveUser = function() {
                LandingService.saveUser(this.user).then(
                    function(response) {
                        $state.go('landing.success');
                    }, 
                    function(response) {
                        if(response.data.error == 'user_duplicate_email') {
                            $state.go('landing.duplicate');
                        }
                        else {
                            
                        }
                    });
            };
            
            this.reload = function() {
                $state.go('landing.form');
            };
    }]);
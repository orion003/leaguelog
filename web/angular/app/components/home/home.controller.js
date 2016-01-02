angular.module('home')
    .controller('HomeSearchController', ['$state', 'HomeService',
        function($state, HomeService){
            ctrl = this;
            HomeService.getLeagues().then(function(result) {
                ctrl.leagues = result.data;
            });
        }])
    .controller('HomeRegisterController', ['$state', 'HomeService', 'AuthService',
        function($state, HomeService, AuthService) {
            function authSuccess(response) {
                console.log(response);
                $("#registration-form").fadeOut("fast", function() {
                    $(this).html('<div class="row"><div class="small-12 columns email-success">Email submitted. Thank you!</div></div>');
                    $(this).fadeIn("slow");
                });
            }

            function authError(response) {
                console.log(response);
                if(response.data.error == "user_duplicate_email") {
                    $("#registration-form").fadeOut("fast", function() {
                        $(this).html('<div class="row"><div class="small-12 columns email-success">Email previously submitted. Thank you!</div></div>');
                        $(this).fadeIn("slow");
                    });
                }
                else {
                    $("#registration-form").fadeOut("fast", function() {
                        $(this).html('<div class="row"><div class="small-12 columns email-fail">There was an error. Please refresh and try again.</div></div>');
                        $(this).fadeIn("slow");
                    });
                }
            }

            this.register = function() {
                AuthService.register(this.user, authSuccess, authError);
            };
        }])
    .controller('HomeLoginController', ['$state', 'HomeService', 'AuthService',
        function($state, HomeService, AuthService) {
            function authSuccess(response) {
                console.log(response);
                $("#registration-form").fadeOut("fast", function() {
                    $(this).html('<div class="row"><div class="small-12 columns email-success">Email submitted. Thank you!</div></div>');
                    $(this).fadeIn("slow");
                });
            }

            function authError(response) {
                console.log(response);
                if(response.data.error == "user_duplicate_email") {
                    $("#registration-form").fadeOut("fast", function() {
                        $(this).html('<div class="row"><div class="small-12 columns email-success">Email previously submitted. Thank you!</div></div>');
                        $(this).fadeIn("slow");
                    });
                }
                else {
                    $("#registration-form").fadeOut("fast", function() {
                        $(this).html('<div class="row"><div class="small-12 columns email-fail">There was an error. Please refresh and try again.</div></div>');
                        $(this).fadeIn("slow");
                    });
                }
            }

            this.login = function() {
                AuthService.login(this.user, authSuccess, authError);
            };
        }]);

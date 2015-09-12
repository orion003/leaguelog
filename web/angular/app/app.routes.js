angular.module('rlApp').config(['$stateProvider', '$urlRouterProvider',
    function($stateProvider, $urlRouterProvider) {
        $stateProvider
            .state('landing', {
                url: '',
                resolve: {
                    user: function() { return {}; }
                },
                templateUrl: '/app/components/landing/landing.html',
                controller: 'LandingController',
                controllerAs: 'landing'
            });
    }]);
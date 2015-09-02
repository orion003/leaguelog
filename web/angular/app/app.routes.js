angular.module('rlApp').config(['$stateProvider', '$urlRouterProvider',
    function($stateProvider, $urlRouterProvider) {
        $urlRouterProvider.otherwise('/');
        
        $stateProvider
            .state('landing', {
                url: '/',
                templateUrl: '/app/components/landing/landing.html',
                controller: 'LandingController',
                controllerAs: 'landing',
                resolve: {
                    users: function() {return {};},
                    user: function() {return {};}
                }
            })
    }]);
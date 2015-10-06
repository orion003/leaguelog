angular.module('league')
    .controller('LeagueStandingsController', ['$state', '$stateParams', 'LeagueService',
        function($state, $stateParams, LeagueService){
            var standingsController = this;
            LeagueService.getStandings($stateParams.leagueId).then(function(result) {
                standingsController.standings = result.data;
            });
        }])
    .controller('LeagueScheduleController', ['$state', '$stateParams', 'LeagueService',
        function($state, $stateParams, LeagueService) {
            var scheduleController = this;
            LeagueService.getSchedules($stateParams.leagueId).then(function(result) {
                scheduleController.schedules = result.data;
            });
        }]);
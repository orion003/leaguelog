angular.module('league')
    .controller('LeagueHeaderController', ['$stateParams', 'LeagueService',
        function($stateParams, LeagueService){
            var headerController = this;
            LeagueService.getLeague($stateParams.leagueId).then(function(result) {
                headerController.league = result.data;
            });
        }])
    .controller('LeagueStandingsController', ['$stateParams', 'LeagueService',
        function($stateParams, LeagueService){
            var standingsController = this;
            LeagueService.getStandings($stateParams.leagueId).then(function(result) {
                standingsController.standings = result.data;
            });
        }])
    .controller('LeagueScheduleController', ['$stateParams', 'LeagueService',
        function($stateParams, LeagueService) {
            var scheduleController = this;
            LeagueService.getSchedules($stateParams.leagueId).then(function(result) {
                scheduleController.schedules = result.data;
            });
        }]);
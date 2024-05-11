package constants

var VenueColumns = []string{"venue_id", "name", "address", "latitude", "longitude", "opening_time", "closing_time", "rating"}

const EventDetailQuery = `
select e.id, e.event_id, e.group_id, e.venue_id, e.created_by, e.name, e.type, e.status,  
 e.params, e.total_cost, e.no_of_participants, e.description, e.start_date_time, 
 e.end_date_time, v.name, v.address, v.latitude, v.longitude 
 from event e join venue v on e.venue_id = v.id `

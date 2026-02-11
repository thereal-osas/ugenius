import { useEffect, useState, useMemo } from "react";
import { Link } from "react-router-dom";
import { Calendar, Clock, MapPin, ArrowRight } from "lucide-react";
import { Button } from "@/components/ui/button";
import { api } from "@/lib/api";
import type  { Event } from "@/lib/api";

// Helper function to create placeholder events (called once outside render)
const createPlaceholderEvents = (): Event[] => {
  const now = Date.now();
  return [
    {
      id: '1',
      title: 'Freshers Breakfast Meeting Unilag',
      description: 'A strategic orientation session for new students focused on building strong academic foundations, mastering university life early, and avoiding common first-year setbacks.',
      type: 'meetup' as const,
      status: 'upcoming' as const,
      start_time: new Date(2026, 1, 14, 9, 0, 0).toISOString(),
      end_time: new Date(2026, 1, 14, 11, 0, 0).toISOString(),
      location: 'Unilag Guest House',
      image_url: 'https://images.unsplash.com/photo-1524178232363-1fb2b075b655?w=800&h=400&fit=crop&q=80',
      is_featured: true,
      created_at: new Date(now).toISOString(),
    },
    {
      id: '2',
      title: 'U-Genius Summit Unilag',
      description: 'A high-impact academic summit bringing together top-performing students to break down the mindset, structure, and discipline required to graduate with excellence at UNILAG.',
      type: 'seminar' as const,
      status: 'upcoming' as const,
      start_time: new Date(2026, 1, 26, 17, 0, 0).toISOString(),
      end_time: new Date(2026, 1, 26, 19, 0, 0).toISOString(),
      location: 'Unilag Guest House',
      image_url: 'https://images.unsplash.com/photo-1541339907198-e08756dedf3f?w=800&h=400&fit=crop&q=80',
      is_featured: true,
      created_at: new Date(now).toISOString(),
    },
    {
      id: '3',
      title: 'Study Group Networking Meetup',
      description: 'Connect with fellow U-Genius members, form study groups, and build lasting academic partnerships.',
      type: 'meetup' as const,
      status: 'upcoming' as const,
      start_time: new Date(now + 21 * 24 * 60 * 60 * 1000).toISOString(),
      end_time: new Date(now + 21 * 24 * 60 * 60 * 1000 + 3 * 60 * 60 * 1000).toISOString(),
      location: 'Online',
      image_url: 'https://images.unsplash.com/photo-1456513080510-7bf3a84b82f8?w=800&h=400&fit=crop&q=80',
      is_featured: true,
      created_at: new Date(now).toISOString(),
    },
  ];
};

const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  return date.toLocaleDateString('en-US', {
    weekday: 'short',
    month: 'short',
    day: 'numeric',
  });
};

const formatTime = (dateString: string) => {
  const date = new Date(dateString);
  return date.toLocaleTimeString('en-US', {
    hour: 'numeric',
    minute: '2-digit',
    hour12: true,
  });
};

const getEventTypeColor = (type: Event['type']) => {
  const colors = {
    workshop: 'bg-blue-100 text-blue-700',
    seminar: 'bg-purple-100 text-purple-700',
    webinar: 'bg-green-100 text-green-700',
    meetup: 'bg-orange-100 text-orange-700',
    contest: 'bg-red-100 text-red-700',
    other: 'bg-gray-100 text-gray-700',
  };
  return colors[type] || colors.other;
};

const EventCard = ({ event }: { event: Event }) => (
  <div className="group bg-card rounded-2xl overflow-hidden shadow-soft hover:shadow-medium transition-all duration-500 hover:-translate-y-2">
    {/* Event Image */}
    <div className="relative h-48 overflow-hidden">
      <img
        src={event.image_url || 'https://images.unsplash.com/photo-1523240795612-9a054b0db644?w=800&h=400&fit=crop&q=80'}
        alt={event.title}
        className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
      />
      <div className="absolute top-4 left-4">
        <span className={`px-3 py-1 rounded-full text-xs font-medium capitalize ${getEventTypeColor(event.type)}`}>
          {event.type}
        </span>
      </div>
    </div>

    {/* Event Content */}
    <div className="p-6">
      <h3 className="font-display text-xl font-semibold text-foreground mb-2 line-clamp-2">
        {event.title}
      </h3>
      <p className="text-muted-foreground text-sm mb-4 line-clamp-2">
        {event.description}
      </p>

      {/* Event Details */}
      <div className="space-y-2 mb-4">
        <div className="flex items-center gap-2 text-sm text-muted-foreground">
          <Calendar className="w-4 h-4 text-gold" />
          <span>{formatDate(event.start_time)}</span>
        </div>
        <div className="flex items-center gap-2 text-sm text-muted-foreground">
          <Clock className="w-4 h-4 text-gold" />
          <span>{formatTime(event.start_time)} - {formatTime(event.end_time)}</span>
        </div>
        <div className="flex items-center gap-2 text-sm text-muted-foreground">
          <MapPin className="w-4 h-4 text-gold" />
          <span>{event.location}</span>
        </div>
      </div>

      {/* CTA */}
      <Button variant="outline" size="sm" className="w-full group/btn">
        Learn More
        <ArrowRight className="w-4 h-4 ml-2 group-hover/btn:translate-x-1 transition-transform" />
      </Button>
    </div>
  </div>
);

const EventsPreview = () => {
  const [events, setEvents] = useState<Event[]>([]);

  // Memoize placeholder events to avoid recreating on each render
  const placeholderEvents = useMemo(() => createPlaceholderEvents(), []);

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        const response = await api.getFeaturedEvents(3);
        if (response.data) {
          setEvents(response.data);
        }
      } catch (error) {
        console.error('Failed to fetch events:', error);
      }
    };

    fetchEvents();
  }, []);

  // Show placeholder events if no events from API
  const displayEvents = events.length > 0 ? events : placeholderEvents;

  return (
    <section className="py-24 bg-background">
      <div className="container mx-auto px-4">
        {/* Section Header */}
        <div className="text-center mb-16">
          <span className="text-gold font-medium text-sm uppercase tracking-wider">
            Upcoming Events
          </span>
          <h2 className="font-display text-4xl md:text-5xl font-bold text-foreground mt-4 mb-6">
            Join Our Community Events
          </h2>
          <p className="text-muted-foreground max-w-2xl mx-auto text-lg">
            Participate in workshops, seminars, and meetups designed to accelerate 
            your academic journey and connect you with like-minded achievers.
          </p>
        </div>

        {/* Events Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {displayEvents.map((event) => (
            <EventCard key={event.id} event={event} />
          ))}
        </div>

        {/* View All CTA */}
        <div className="text-center mt-12">
          <Link to="/join">
            <Button variant="hero" size="lg" className="group">
              Join to Access All Events
              <ArrowRight className="w-5 h-5 ml-2 group-hover:translate-x-1 transition-transform" />
            </Button>
          </Link>
        </div>
      </div>
    </section>
  );
};

export default EventsPreview;


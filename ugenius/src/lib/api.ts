// API Client for U-Genius Backend

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1';

interface ApiResponse<T = unknown> {
  success: boolean;
  message: string;
  data?: T;
  error?: string;
  meta?: {
    page: number;
    page_size: number;
    total: number;
    total_pages: number;
  };
}

interface RequestOptions {
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE';
  body?: unknown;
  headers?: Record<string, string>;
}

class ApiClient {
  private baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  private getToken(): string | null {
    return localStorage.getItem('access_token');
  }

  private async request<T>(endpoint: string, options: RequestOptions = {}): Promise<ApiResponse<T>> {
    const { method = 'GET', body, headers = {} } = options;

    const token = this.getToken();
    const requestHeaders: Record<string, string> = {
      'Content-Type': 'application/json',
      ...headers,
    };

    if (token) {
      requestHeaders['Authorization'] = `Bearer ${token}`;
    }

    const config: RequestInit = {
      method,
      headers: requestHeaders,
    };

    if (body) {
      config.body = JSON.stringify(body);
    }

    const response = await fetch(`${this.baseUrl}${endpoint}`, config);
    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || data.message || 'An error occurred');
    }

    return data;
  }

  // Auth endpoints
  async register(data: {
    email: string;
    password: string;
    first_name: string;
    last_name: string;
    campus_id?: string;
  }) {
    return this.request('/auth/register', { method: 'POST', body: data });
  }

  async login(email: string, password: string) {
    return this.request<{ access_token: string; refresh_token: string; user: unknown }>(
      '/auth/login',
      { method: 'POST', body: { email, password } }
    );
  }

  async refreshToken(refreshToken: string) {
    return this.request<{ access_token: string; refresh_token: string }>(
      '/auth/refresh',
      { method: 'POST', body: { refresh_token: refreshToken } }
    );
  }

  async logout() {
    return this.request('/auth/logout', { method: 'POST' });
  }

  async getMe() {
    return this.request<User>('/me');
  }

  async updateProfile(data: Partial<User>) {
    return this.request<User>('/me', { method: 'PUT', body: data });
  }

  // Campus endpoints
  async getCampuses() {
    return this.request<Campus[]>('/campuses');
  }

  async getCampus(id: string) {
    return this.request<Campus>(`/campuses/${id}`);
  }

  // Reading Hours endpoints
  async getReadingHours(params?: { page?: number; status?: string }) {
    const query = new URLSearchParams();
    if (params?.page) query.set('page', params.page.toString());
    if (params?.status) query.set('status', params.status);
    return this.request<ReadingHour[]>(`/reading-hours?${query}`);
  }

  async createReadingHour(data: {
    subject: string;
    duration_minutes: number;
    reading_date: string;
    notes?: string;
  }) {
    return this.request<ReadingHour>('/reading-hours', { method: 'POST', body: data });
  }

  async getReadingHourStats() {
    return this.request<ReadingHourStats>('/reading-hours/stats');
  }

  // Goals endpoints
  async getGoals() {
    return this.request<Goal[]>('/goals');
  }

  async createGoal(data: { title: string; description?: string; target_value: number; goal_type: string }) {
    return this.request<Goal>('/goals', { method: 'POST', body: data });
  }

  // Achievements endpoints
  async getAchievements() {
    return this.request<Achievement[]>('/achievements');
  }

  async getAllBadges() {
    return this.request<Badge[]>('/achievements/badges');
  }

  // Leaderboard
  async getLeaderboard(campusId?: string) {
    const query = campusId ? `?campus_id=${campusId}` : '';
    return this.request<LeaderboardEntry[]>(`/leaderboard${query}`);
  }

  // Notifications
  async getNotifications() {
    return this.request<Notification[]>('/notifications');
  }

  async getUnreadCount() {
    return this.request<{ count: number }>('/notifications/unread-count');
  }

  async markNotificationRead(id: string) {
    return this.request(`/notifications/${id}/read`, { method: 'POST' });
  }

  // Events
  async getEvents(page = 1, pageSize = 10) {
    return this.request<Event[]>(`/events?page=${page}&page_size=${pageSize}`);
  }

  async getFeaturedEvents(limit = 3) {
    return this.request<Event[]>(`/events/featured?limit=${limit}`);
  }

  async getEventById(id: string) {
    return this.request<Event>(`/events/${id}`);
  }
}

export const api = new ApiClient(API_BASE_URL);

// Type definitions
export interface User {
  id: string;
  email: string;
  first_name: string;
  last_name: string;
  avatar_url?: string;
  role: 'student' | 'campus_admin' | 'super_admin';
  campus_id?: string;
  is_verified: boolean;
}

export interface Campus {
  id: string;
  name: string;
  code: string;
  location: string;
  state: string;
  country: string;
  is_active: boolean;
}

export interface ReadingHour {
  id: string;
  user_id: string;
  subject: string;
  duration_minutes: number;
  reading_date: string;
  notes?: string;
  status: 'pending' | 'approved' | 'rejected';
  created_at: string;
}

export interface ReadingHourStats {
  total_hours: number;
  weekly_hours: number;
  total_submissions: number;
  approved_submissions: number;
  pending_submissions: number;
}

export interface Goal {
  id: string;
  title: string;
  description?: string;
  goal_type: string;
  target_value: number;
  current_value: number;
  is_completed: boolean;
  start_date: string;
  end_date?: string;
}

export interface Achievement {
  id: string;
  badge_type: string;
  badge_name: string;
  badge_description: string;
  earned_at: string;
}

export interface Badge {
  type: string;
  name: string;
  description: string;
  icon: string;
}

export interface LeaderboardEntry {
  rank: number;
  user_id: string;
  first_name: string;
  last_name: string;
  avatar_url?: string;
  campus_name: string;
  total_hours: number;
  submission_count: number;
}

export interface Notification {
  id: string;
  type: string;
  title: string;
  message: string;
  is_read: boolean;
  created_at: string;
}

export interface Event {
  id: string;
  title: string;
  description: string;
  type: 'workshop' | 'seminar' | 'webinar' | 'meetup' | 'contest' | 'other';
  status: 'upcoming' | 'ongoing' | 'completed' | 'cancelled';
  start_time: string;
  end_time: string;
  location: string;
  virtual_link?: string;
  image_url?: string;
  campus_id?: string;
  max_attendees?: number;
  is_featured: boolean;
  created_at: string;
}


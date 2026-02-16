import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useToast } from "@/hooks/use-toast";
import { CheckCircle, Users, BookOpen, Award } from "lucide-react";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080/api/v1";

interface Campus {
  id: string;
  name: string;
  code: string;
}

const benefits = [
  { icon: Users, text: "Join 500+ ambitious students" },
  { icon: BookOpen, text: "Access premium study materials" },
  { icon: Award, text: "Scholarship opportunities" },
];

const levels = [
  { value: "100", label: "100 Level" },
  { value: "200", label: "200 Level" },
  { value: "300", label: "300 Level" },
  { value: "400", label: "400 Level" },
  { value: "500", label: "500 Level" },
  { value: "600", label: "600 Level" },
  { value: "postgraduate", label: "Postgraduate" },
];

const Join = () => {
  const { toast } = useToast();
  const navigate = useNavigate();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [campuses, setCampuses] = useState<Campus[]>([]);
  const [formData, setFormData] = useState({
    first_name: "",
    last_name: "",
    email: "",
    phone: "",
    campus_id: "",
    level: "",
    faculty: "",
    department: "",
    address: "",
  });

  // Fetch campuses on mount
  useEffect(() => {
    const fetchCampuses = async () => {
      try {
        const response = await fetch(`${API_URL}/campuses`);
        if (response.ok) {
          const data = await response.json();
          setCampuses(data.data || []);
        }
      } catch (error) {
        console.error("Failed to fetch campuses:", error);
      }
    };
    fetchCampuses();
  }, []);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    setFormData((prev) => ({
      ...prev,
      [e.target.name]: e.target.value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);

    try {
      const response = await fetch(`${API_URL}/auth/register`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: formData.email,
          password: "defaultPassword123!", // Default password for students
          first_name: formData.first_name,
          last_name: formData.last_name,
          phone: formData.phone,
          campus_id: formData.campus_id || undefined,
          institution: "U-Genius Platform",
          department: formData.department,
          level: formData.level,
        }),
      });

      const data = await response.json();

      if (response.ok) {
        // Navigate to thank you page on success
        navigate("/thank-you");
      } else if (response.status === 409) {
        toast({
          title: "Already Registered",
          description: data.message || "This email is already registered.",
          variant: "destructive",
        });
      } else {
        toast({
          title: "Registration Failed",
          description: data.message || "Something went wrong. Please try again.",
          variant: "destructive",
        });
      }
    } catch (error) {
      toast({
        title: "Connection Error",
        description: "Unable to connect to the server. Please try again later.",
        variant: "destructive",
      });
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="min-h-screen">
      <Navbar />
      <main>
        <section className="pt-32 pb-24 bg-gradient-hero relative overflow-hidden">
          <div className="absolute inset-0 opacity-30">
            <div className="absolute top-1/4 left-1/4 w-96 h-96 bg-gold/20 rounded-full blur-3xl" />
            <div className="absolute bottom-1/4 right-1/4 w-96 h-96 bg-coral/20 rounded-full blur-3xl" />
          </div>

          <div className="container mx-auto px-4 relative z-10">
            <div className="max-w-6xl mx-auto">
              <div className="grid grid-cols-1 lg:grid-cols-2 gap-16 items-start">
                {/* Left Column - Info */}
                <div className="flex flex-col justify-center">
                  <span className="text-gold font-medium text-sm uppercase tracking-wider">
                    Join U-Genius
                  </span>
                  <h1 className="font-display text-5xl md:text-6xl font-bold text-foreground mt-4 mb-6">
                    Start Your Journey to Excellence
                  </h1>
                  <p className="text-lg text-muted-foreground leading-relaxed mb-8">
                    Take the first step towards academic greatness. Fill out the form 
                    to become a member of U-Genius and unlock your full potential.
                  </p>

                  <div className="space-y-4">
                    {benefits.map((benefit) => (
                      <div key={benefit.text} className="flex items-center gap-3">
                        <div className="w-10 h-10 rounded-full bg-gold/10 flex items-center justify-center">
                          <benefit.icon className="w-5 h-5 text-gold" />
                        </div>
                        <span className="text-foreground font-medium">{benefit.text}</span>
                      </div>
                    ))}
                  </div>
                </div>

                {/* Right Column - Form */}
                <div className="bg-card rounded-3xl p-8 md:p-10 shadow-medium">
                  <h2 className="font-display text-2xl font-semibold text-foreground mb-6">
                    Registration Form
                  </h2>

                  <form onSubmit={handleSubmit} className="space-y-5">
                    {/* First Name */}
                    <div className="space-y-2">
                      <Label htmlFor="first_name">First Name *</Label>
                      <Input
                        id="first_name"
                        name="first_name"
                        value={formData.first_name}
                        onChange={handleChange}
                        required
                        placeholder="John"
                        className="bg-background"
                      />
                    </div>

                    {/* Last Name */}
                    <div className="space-y-2">
                      <Label htmlFor="last_name">Last Name *</Label>
                      <Input
                        id="last_name"
                        name="last_name"
                        value={formData.last_name}
                        onChange={handleChange}
                        required
                        placeholder="Doe"
                        className="bg-background"
                      />
                    </div>

                    {/* Email */}
                    <div className="space-y-2">
                      <Label htmlFor="email">Email Address *</Label>
                      <Input
                        id="email"
                        name="email"
                        type="email"
                        value={formData.email}
                        onChange={handleChange}
                        required
                        placeholder="john@university.edu"
                        className="bg-background"
                      />
                    </div>

                    {/* Phone */}
                    <div className="space-y-2">
                      <Label htmlFor="phone">Phone Number</Label>
                      <Input
                        id="phone"
                        name="phone"
                        type="tel"
                        value={formData.phone}
                        onChange={handleChange}
                        placeholder="+234 800 000 0000"
                        className="bg-background"
                      />
                    </div>

                    {/* University/Campus */}
                    <div className="space-y-2">
                      <Label htmlFor="campus_id">University/Campus</Label>
                      <select
                        id="campus_id"
                        name="campus_id"
                        value={formData.campus_id}
                        onChange={handleChange}
                        className="w-full h-10 px-3 rounded-lg border border-input bg-background text-foreground text-sm focus:outline-none focus:ring-2 focus:ring-ring"
                      >
                        <option value="">Select your university</option>
                        {campuses.map((campus) => (
                          <option key={campus.id} value={campus.id}>
                            {campus.name}
                          </option>
                        ))}
                      </select>
                    </div>

                    {/* Level */}
                    <div className="space-y-2">
                      <Label htmlFor="level">Level</Label>
                      <select
                        id="level"
                        name="level"
                        value={formData.level}
                        onChange={handleChange}
                        className="w-full h-10 px-3 rounded-lg border border-input bg-background text-foreground text-sm focus:outline-none focus:ring-2 focus:ring-ring"
                      >
                        <option value="">Select your level</option>
                        {levels.map((level) => (
                          <option key={level.value} value={level.value}>
                            {level.label}
                          </option>
                        ))}
                      </select>
                    </div>

                    {/* Faculty & Department */}
                    <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                      <div className="space-y-2">
                        <Label htmlFor="faculty">Faculty/School</Label>
                        <Input
                          id="faculty"
                          name="faculty"
                          value={formData.faculty}
                          onChange={handleChange}
                          placeholder="e.g. Engineering"
                          className="bg-background"
                        />
                      </div>
                      <div className="space-y-2">
                        <Label htmlFor="department">Department</Label>
                        <Input
                          id="department"
                          name="department"
                          value={formData.department}
                          onChange={handleChange}
                          placeholder="e.g. Computer Science"
                          className="bg-background"
                        />
                      </div>
                    </div>

                    {/* Address in School */}
                    <div className="space-y-2">
                      <Label htmlFor="address">Address in School</Label>
                      <Input
                        id="address"
                        name="address"
                        value={formData.address}
                        onChange={handleChange}
                        placeholder="e.g. Room 12, Block A, Jaja Hall"
                        className="bg-background"
                      />
                    </div>

                    <Button
                      type="submit"
                      variant="hero"
                      size="xl"
                      className="w-full"
                      disabled={isSubmitting}
                    >
                      {isSubmitting ? (
                        <span className="flex items-center gap-2">
                          <div className="w-5 h-5 border-2 border-primary border-t-transparent rounded-full animate-spin" />
                          Submitting...
                        </span>
                      ) : (
                        <span className="flex items-center gap-2">
                          <CheckCircle className="w-5 h-5" />
                          Submit Application
                        </span>
                      )}
                    </Button>
                  </form>
                </div>
              </div>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </div>
  );
};

export default Join;

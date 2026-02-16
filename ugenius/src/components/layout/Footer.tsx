import { Link } from "react-router-dom";
import {  Mail, MapPin, Phone } from "lucide-react";
import logo from "@/assets/ugenius_logo.png"

const Footer = () => {
  const currentYear = new Date().getFullYear();

  return (
    <footer className="bg-primary text-primary-foreground">
      <div className="container mx-auto px-4 py-16">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-12">
          {/* Brand */}
          <div className="space-y-4">
            <Link to="/">
              <div className="flex items-center gap-2">
                <div className="flex items-center justify-center">
                  <img 
                    src={logo} 
                    alt="U-Genius Logo" 
                    className="w-full h-full object-contain rounded-full"
                  />
                </div>
              </div>
            </Link>
            <p className="text-primary-foreground/80 text-sm leading-relaxed">
              Empowering students to achieve academic excellence through the transformative Auto-Bio Jacking system.
            </p>
           
          </div>

          {/* Quick Links */}
          <div>
            <h4 className="font-display text-lg font-semibold mb-4">Quick Links</h4>
            <ul className="space-y-3">
              {[
                { name: "About Us", path: "/about" },
                { name: "Benefits", path: "/benefits" },
                { name: "Resources", path: "/resources" },
                { name: "Testimonials", path: "/testimonials" },
                { name: "Join Us", path: "/join" },
              ].map((link) => (
                <li key={link.path}>
                  <Link
                    to={link.path}
                    className="text-primary-foreground/70 hover:text-gold transition-colors duration-300 text-sm"
                  >
                    {link.name}
                  </Link>
                </li>
              ))}
            </ul>
          </div>

          {/* Resources */}
          <div>
            <h4 className="font-display text-lg font-semibold mb-4">Resources</h4>
            <ul className="space-y-3">
              {[
                { name: "Study Materials", path: "/resources" },
                { name: "Video Tutorials", path: "/resources" },
                { name: "PDF Guides", path: "/resources" },
                { name: "Mentorship Program", path: "/benefits" },
                { name: "FAQs", path: "/contact" },
              ].map((link) => (
                <li key={link.name}>
                  <Link
                    to={link.path}
                    className="text-primary-foreground/70 hover:text-gold transition-colors duration-300 text-sm"
                  >
                    {link.name}
                  </Link>
                </li>
              ))}
            </ul>
          </div>

          {/* Contact */}
          <div>
            <h4 className="font-display text-lg font-semibold mb-4">Contact Us</h4>
            <ul className="space-y-4">
              <li className="flex items-start gap-3">
                <MapPin size={18} className="text-gold mt-0.5 shrink-0" />
                <span className="text-primary-foreground/70 text-sm">
                  Plot 11A-C Kudirat Abiola Way Alausa <br />Ikeja, Lagos, Nigeria                </span>
              </li>
              <li className="flex items-center gap-3">
                <Mail size={18} className="text-gold shrink-0" />
                <a
                  href="mailto:ugeniusclub@gmail.com"
                  className="text-primary-foreground/70 hover:text-gold transition-colors duration-300 text-sm"
                >
                  ugeniusclublagos@gmail.com
                </a>
              </li>
              <li className="flex items-center gap-3">
                <Phone size={18} className="text-gold shrink-0" />
                <a
                  href="tel:+2347062557356"
                  className="text-primary-foreground/70 hover:text-gold transition-colors duration-300 text-sm"
                >
                  +234 706 255 7356
                </a>
              </li>
            </ul>
          </div>
        </div>

        {/* Bottom */}
        <div className="mt-12 pt-8 border-t border-primary-foreground/10 flex flex-col md:flex-row justify-between items-center gap-4">
          <p className="text-primary-foreground/60 text-sm">
            © {currentYear} U-Genius Academic Club. All rights reserved.
          </p>
          <div className="flex gap-6">
            <a href="#" className="text-primary-foreground/60 hover:text-gold transition-colors text-sm">
              Privacy Policy
            </a>
            <a href="#" className="text-primary-foreground/60 hover:text-gold transition-colors text-sm">
              Terms of Service
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;

import React from 'react';

const Footer = () => {
    return (
    <footer className="bg-gray-800 text-white py-6">
      <div className="max-w-screen-xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {/* About Section */}
          <div>
            <h3 className="text-lg font-semibold mb-2">О нас</h3>
            <p className="text-gray-400 text-sm">
              EduConnect — это платформа для студентов, преподавателей и работодателей, которая улучшает образовательный процесс и помогает найти стажировки и вакансии.
            </p>
          </div>

          {/* Quick Links */}
          <div>
            <h3 className="text-lg font-semibold mb-2">Ссылки</h3>
            <ul className="text-gray-400 text-sm">
              <li>
                <a href="/" className="hover:underline">Главная</a>
              </li>
              <li>
                <a href="/courses" className="hover:underline">Курсы</a>
              </li>
              <li>
                <a href="/jobs" className="hover:underline">Вакансии</a>
              </li>
              <li>
                <a href="/forum" className="hover:underline">Форум</a>
              </li>
            </ul>
          </div>

          {/* Contact Section */}
          <div>
            <h3 className="text-lg font-semibold mb-2">Контакты</h3>
            <ul className="text-gray-400 text-sm">
              <li>
                <p>Email: <a href="mailto:support@educonnect.com" className="hover:underline">support@educonnect.com</a></p>
              </li>
              <li>
                <p>Телефон: <a href="tel:+123456789" className="hover:underline">+1 234 567 89</a></p>
              </li>
            </ul>
          </div>

          {/* Social Media Links */}
          <div>
            <h3 className="text-lg font-semibold mb-2">Следите за нами</h3>
            <div className="flex space-x-4">
              <a href="https://www.facebook.com" target="_blank" rel="noopener noreferrer">
                <img src="/facebook-icon.svg" alt="Facebook" className="w-6 h-6 text-gray-400 hover:text-white" />
              </a>
              <a href="https://www.twitter.com" target="_blank" rel="noopener noreferrer">
                <img src="/twitter-icon.svg" alt="Twitter" className="w-6 h-6 text-gray-400 hover:text-white" />
              </a>
              <a href="https://www.linkedin.com" target="_blank" rel="noopener noreferrer">
                <img src="/linkedin-icon.svg" alt="LinkedIn" className="w-6 h-6 text-gray-400 hover:text-white" />
              </a>
              <a href="https://www.instagram.com" target="_blank" rel="noopener noreferrer">
                <img src="/instagram-icon.svg" alt="Instagram" className="w-6 h-6 text-gray-400 hover:text-white" />
              </a>
            </div>
          </div>
        </div>

        <div className="mt-8 border-t border-gray-700 pt-4 text-center">
          <p className="text-gray-400 text-sm">
            &copy; {new Date().getFullYear()} EduConnect. Все права защищены.
          </p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
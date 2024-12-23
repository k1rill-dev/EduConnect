import React from 'react';
import { Navbar, Dropdown, Button } from 'flowbite-react';

const Header = () => {
  return (
    <Navbar fluid={true} rounded={true} className={"mr-5"}>
      <Navbar.Brand href="/">
        <img
          src="/logo.svg"
          className="mr-3 h-6 sm:h-9"
          alt="EduConnect Logo"
        />
        <span className="self-center whitespace-nowrap text-xl font-semibold text-gray-800 dark:text-white">
          EduConnect
        </span>
      </Navbar.Brand>
      <Navbar.Toggle />
      <Navbar.Collapse>
        <Navbar.Link href="/" active={true}>
          Главная
        </Navbar.Link>
        <Navbar.Link href="/courses">Курсы</Navbar.Link>
        <Navbar.Link href="/jobs">Вакансии</Navbar.Link>
        <Navbar.Link href="/forum">Форум</Navbar.Link>

        {/* Выпадающий список "Профиль" */}
        <Dropdown
          label="Профиль"
          inline={true}
          arrowIcon={true}
          className="relative z-10"
        >
          <Dropdown.Item href="/profile">
            Мой профиль
          </Dropdown.Item>
          <Dropdown.Item href="/settings">
            Настройки
          </Dropdown.Item>
          <Dropdown.Divider />
          <Dropdown.Item onClick={() => alert('Вышли из аккаунта')}>
            Выйти
          </Dropdown.Item>
        </Dropdown>
      </Navbar.Collapse>
    </Navbar>
  );
};

export default Header;

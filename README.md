# Wikipedia Transport Hub Crawler – Database Documentation

## Overview

This database stores transport hubs and their connections for **airports** and **train stations** (including metro, regional, and high-speed rail).
It supports:

* Hubs (stations/airports) with GPS coordinates
* Operators (airlines, rail operators, metro authorities)
* Multi-stop services (e.g., Boston → NYC → DC)
* Routes for air, rail, and metro

## **1. Transport\_Hubs**

Stores all hubs — airports, train stations, and metro stops.

| Field       | Type         | Description                                               |
| ----------- | ------------ | --------------------------------------------------------- |
| `hub_id`    | INTEGER      | Unique identifier                                         |
| `name`      | TEXT         | Official name of the hub                                  |
| `type`      | TEXT         | Type of hub (`airport`, `train_station`, `metro_station`) |
| `iata_code` | CHAR(3)      | IATA code (NULL if not applicable)                        |
| `icao_code` | CHAR(4)      | ICAO code (NULL if not applicable)                        |
| `city`      | TEXT         | City where the hub is located                             |
| `country`   | TEXT         | Country where the hub is located                          |
| `latitude`  | DECIMAL(9,6) | GPS latitude                                              |
| `longitude` | DECIMAL(9,6) | GPS longitude                                             |
| `wiki_url`  | TEXT         | Wikipedia page for the hub                                |

---

## **2. Operators**

Represents companies running transport services.

| Field         | Type    | Description                                |
| ------------- | ------- | ------------------------------------------ |
| `operator_id` | INTEGER | Unique identifier                          |
| `name`        | TEXT    | Operator name                              |
| `type`        | TEXT    | `airline`, `rail`, `metro`                 |
| `iata_code`   | CHAR(2) | Airline IATA code (NULL if not applicable) |
| `icao_code`   | CHAR(3) | Airline ICAO code (NULL if not applicable) |
| `country`     | TEXT    | Country of origin                          |
| `wiki_url`    | TEXT    | Wikipedia page for the operator            |

---

## **3. Service\_Lines**

Defines a continuous service (e.g., a specific train line or flight number).

| Field         | Type    | Description                                |
| ------------- | ------- | ------------------------------------------ |
| `service_id`  | INTEGER | Unique identifier                          |
| `operator_id` | INTEGER | References `Operators.operator_id`         |
| `mode`        | TEXT    | `air`, `rail`, `metro`                     |
| `name`        | TEXT    | Service name or flight number              |
| `seasonal`    | BOOLEAN | If the service is seasonal                 |
| `notes`       | TEXT    | Additional info (e.g., suspended, express) |

---

## **4. Service\_Stops**

Lists all stops for a service in travel order.

| Field            | Type    | Description                                 |
| ---------------- | ------- | ------------------------------------------- |
| `service_id`     | INTEGER | References `Service_Lines.service_id`       |
| `stop_order`     | INTEGER | Order of stop in route (1 = starting point) |
| `hub_id`         | INTEGER | References `Transport_Hubs.hub_id`          |


---

## **5. Routes **

If you need direct origin–destination pairs for mapping or analysis, this can be **generated** from `Service_Stops`.

| Field            | Type    | Description                           |
| ---------------- | ------- | ------------------------------------- |
| `origin_id`      | INTEGER | References `Transport_Hubs.hub_id`    |
| `destination_id` | INTEGER | References `Transport_Hubs.hub_id`    |
| `service_id`     | INTEGER | References `Service_Lines.service_id` |
| `operator_id`    | INTEGER | References `Operators.operator_id`    |
| `mode`           | TEXT    | `air`, `rail`, `metro`                |

---

## Relationships

* **Transport\_Hubs** stores all stations/airports.
* **Operators** are linked to **Service\_Lines** via `operator_id`.
* **Service\_Lines** are linked to **Service\_Stops** to record the full stop sequence.
* **Routes** can be derived for direct connections between stops.

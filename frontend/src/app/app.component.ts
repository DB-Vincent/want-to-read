import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { ApiService } from './services/api.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {
  healthStatus: string = 'Checking...';
  isHealthy: boolean = false;

  constructor(private apiService: ApiService) {}

  ngOnInit() {
    this.checkHealth();
  }

  private checkHealth() {
    this.apiService.checkHealth().subscribe({
      next: (response) => {
        this.healthStatus = response.status;
        this.isHealthy = response.status === 'ok';
      },
      error: (error) => {
        this.healthStatus = 'Error connecting to backend';
        this.isHealthy = false;
        console.error('Health check failed:', error);
      }
    });
  }
}

import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { TypeOrmModule } from '@nestjs/typeorm';
import { UsersModule } from './modules/users/users.module';
import { AuthModule } from './modules/auth/auth.module';
import { ActivitiesModule } from './modules/activities/activities.module';
import { ChallengesModule } from './modules/challenges/challenges.module';
import { AchievementsModule } from './modules/achievements/achievements.module';
import { MapsModule } from './modules/maps/maps.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    TypeOrmModule.forRoot({
      type: 'postgres',
      host: process.env.DB_HOST || 'localhost',
      port: parseInt(process.env.DB_PORT, 10) || 5432,
      username: process.env.DB_USERNAME || 'postgres',
      password: process.env.DB_PASSWORD || 'postgres',
      database: process.env.DB_DATABASE || 'push_run_app',
      entities: [__dirname + '/**/*.entity{.ts,.js}'],
      synchronize: process.env.NODE_ENV !== 'production',
    }),
    UsersModule,
    AuthModule,
    ActivitiesModule,
    ChallengesModule,
    AchievementsModule,
    MapsModule,
  ],
})
export class AppModule {} 